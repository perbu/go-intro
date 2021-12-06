package main

import (
	_ "embed"
	"fmt"
	"github.com/gliderlabs/ssh"
	gossh "golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"log"
)

//go:embed id_rsa.pub
var embeddedAuthorizedKeys []byte

//go:embed privkey
var embeddedPrivateKey []byte

type sshApp struct {
	server        *ssh.Server
	pubKeyHandler func(ctx ssh.Context, key ssh.PublicKey) bool
	pubKey        ssh.PublicKey
}

func getAuthorizedKeys() gossh.PublicKey {
	pubKey, _, _, _, err := gossh.ParseAuthorizedKey(embeddedAuthorizedKeys)
	if err != nil {
		log.Fatalf("Couldn't make sense of authorized key: %s", err)
	}
	return pubKey
}
func getPrivateKey() gossh.Signer {
	pubKey, err := gossh.ParseRawPrivateKey(embeddedPrivateKey)
	if err != nil {
		log.Fatalf("Couldn't make sense of embedded private key: %s", err)
	}
	signer, err := gossh.NewSignerFromKey(pubKey)
	if err != nil {
		log.Fatalf("Couldn't make sense of private key: %s", err)
	}
	return signer
}

func sshHandler(s ssh.Session) {
	defer func(s ssh.Session) { // close the connection on method return. log error if any.
		err := s.Close()
		if err != nil {
			fmt.Printf("Error closing connection: %s", err)
		}
	}(s)
	if s.RawCommand() != "" {
		io.WriteString(s, "raw commands are not supported")
		return
	}

	term := terminal.NewTerminal(s, fmt.Sprintf("%s> ", s.User()))
	pty, winCh, isPty := s.Pty()
	if isPty {
		fmt.Println("PTY term", pty.Term)
		go func() {
			for chInfo := range winCh {
				fmt.Println("winch:", chInfo)
				err := term.SetSize(chInfo.Width, chInfo.Height)
				if err != nil {
					fmt.Println("winch error:", err)
				}
			}
		}()
	}

	for {
		line, err := term.ReadLine()
		if err == io.EOF {
			// Ignore errors here:
			_, _ = io.WriteString(s, "EOF.\n")
			break
		}
		if err != nil {
			// Ignore errors here:
			_, _ = io.WriteString(s, "Error while reading: "+err.Error())
			break
		}
		if line == "quit" {
			break
		}
		if line == "" {
			continue
		}
		output, err := handleTerminalInput(line)
		if err != nil {
			log.Printf("Error handling terminal input: %s", err)
			return
		}
		io.WriteString(s, output)
	}
	io.WriteString(s, fmt.Sprintf("Welcome to my own ssh daemon, %s\n", s.User()))
}

func handleTerminalInput(line string) (string, error) {
	return fmt.Sprintf("echo: %s\n", line), nil
}

func (a sshApp) myPubKeyHandler(ctx ssh.Context, key ssh.PublicKey) bool {
	if ssh.KeysEqual(key, a.pubKey) {
		return true
	} else {
		return false
	}
}

func main() {
	authorizedKey := getAuthorizedKeys()
	serverKey := getPrivateKey()
	port := 2000
	server := &ssh.Server{
		LocalPortForwardingCallback: nil,
		Addr:                        fmt.Sprintf(":%d", port),
		Handler:                     sshHandler,
	}
	server.AddHostKey(serverKey)
	app := sshApp{
		server: server,
		pubKey: authorizedKey,
	}
	app.pubKeyHandler = app.myPubKeyHandler
	log.Printf("starting ssh server on port %d...", port)
	server.ListenAndServe()
}
