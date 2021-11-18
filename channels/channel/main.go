package channel

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/perbu/go-intro/channels/bank"
)

type bankMessageType int

const (
	returnMessage bankMessageType = iota
	createAccount
	getBalance
	deposit
	withdraw
	transfer
	quit
)

type channelBank struct {
	name     string
	accounts accountMap
	cChan    chan bankMessage
}

type bankMessage struct {
	messageType   bankMessageType
	account       uuid.UUID
	targetAccount uuid.UUID
	amount        int64
	rChan         chan bankMessage
}

type account struct {
	balance int64
}
type accountMap map[uuid.UUID]*account

func CreateBank(name string) bank.Bank {
	c := channelBank{
		name:     name,
		accounts: make(accountMap),
		cChan:    make(chan bankMessage),
	}
	go c.banker()
	return &c
}

func (b *channelBank) banker() {
	for {
		msg := <-b.cChan
		switch msg.messageType {
		case createAccount:
			id, _ := uuid.NewV4()
			b.accounts[*id] = &account{}
			ret := bankMessage{messageType: returnMessage, account: *id}
			msg.rChan <- ret
		case getBalance:
			a := b.getAccount(msg.account)
			ret := bankMessage{amount: a.balance}
			msg.rChan <- ret
		case deposit:
			a := b.getAccount(msg.account)
			a.balance = a.balance + msg.amount
			ret := bankMessage{messageType: returnMessage, amount: a.balance}
			msg.rChan <- ret
		case withdraw:
			a := b.getAccount(msg.account)
			a.balance = a.balance - msg.amount
			ret := bankMessage{messageType: returnMessage, amount: a.balance}
			msg.rChan <- ret
		case transfer:
			from := b.getAccount(msg.account)
			to := b.getAccount(msg.account)
			from.balance = from.balance - msg.amount
			to.balance = to.balance + msg.amount
			ret  := bankMessage{messageType: returnMessage, amount: msg.amount}
			msg.rChan <- ret
		case quit:
			close(b.cChan)
			fmt.Println("Banker goroutine shutting down")
			break
		}
	}
}

func (b *channelBank) CreateAccount() uuid.UUID {
	msg := bankMessage{
		messageType: createAccount,
		rChan:       make(chan bankMessage),
	}
	b.cChan <- msg
	ret := <-msg.rChan
	close(msg.rChan)
	return ret.account
}

func (b *channelBank) Balance(account uuid.UUID) int64 {
	msg := bankMessage{
		messageType: getBalance,
		account:     account,
		rChan:       make(chan bankMessage),
	}
	b.cChan <- msg
	ret := <-msg.rChan
	close(msg.rChan)
	return ret.amount
}

func (b *channelBank) getAccount(account uuid.UUID) *account {
	return b.accounts[account]
}

func (b *channelBank) Deposit(account uuid.UUID, amount int64) int64 {
	msg := bankMessage{
		messageType: deposit,
		amount:      amount,
		account:     account,
		rChan:       make(chan bankMessage),
	}
	b.cChan <- msg
	ret := <-msg.rChan
	close(msg.rChan)
	return ret.amount
}

func (b *channelBank) Withdraw(account uuid.UUID, amount int64) int64 {
	msg := bankMessage{
		messageType: withdraw,
		account:     account,
		amount:      amount,
		rChan:       make(chan bankMessage),
	}
	b.cChan <- msg
	ret := <-msg.rChan
	close(msg.rChan)
	return ret.amount
}
func (b *channelBank) Transfer(from, to uuid.UUID, amount int64 ) {
	msg := bankMessage{
		messageType:   transfer,
		account:       from,
		targetAccount: to,
		amount:        amount,
		rChan:         make(chan bankMessage),
	}
	b.cChan <- msg
	close(msg.rChan)
	return
}


func (b *channelBank) Close() error {
	msg := bankMessage{messageType: quit}
	b.cChan <- msg
	return nil
}
