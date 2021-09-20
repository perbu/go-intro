package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type customError struct {
	code      int
	location  string
	goroutine int
	secret    string
}

func (cE customError) Error() string {
	return fmt.Sprintf("Error %d in %s [goroutine %d]",
		cE.code, cE.location, cE.goroutine)
}

func doFunkyStuff() error {
	err := customError{
		code:      512,
		location:  "Oslo",
		goroutine: goid(),
		secret: "geheim",
	}
	return err
}

func main() {
	err := doFunkyStuff()
	if err != nil {
		fmt.Printf("Got error: %s\n", err)
		cErr, ok := err.(customError)
		if ok {
			fmt.Printf("I've found the secret: %s\n", cErr.secret)
		} else {
			fmt.Println("Cast failed.")
		}

	}

}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
