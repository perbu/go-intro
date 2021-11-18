package naive

import (
	uuid "github.com/nu7hatch/gouuid"
	"github.com/perbu/go-intro/channels/bank"
)

type account struct {
	balance int64
}

type accountMap map[uuid.UUID]*account

type classicBank struct {
	name     string
	accounts accountMap
}

func CreateBank(name string) bank.Bank {
	c := classicBank{
		name:     name,
		accounts: make(accountMap),
	}
	return &c
}

func (b *classicBank) CreateAccount() uuid.UUID {
	accId, _ := uuid.NewV4()
	acc := &account{
		balance: 0,
	}
	b.accounts[*accId] = acc
	return *accId
}

func (b *classicBank) getAccount(account uuid.UUID) *account {
	acct, _ := b.accounts[account]
	return acct
}

func (b *classicBank) Deposit(account uuid.UUID, amount int64) int64 {
	a := b.getAccount(account)
	a.balance = a.balance + amount
	return a.balance
}

func (b *classicBank) Withdraw(account uuid.UUID, amount int64) int64 {
	a := b.getAccount(account)
	a.balance = a.balance - amount
	return a.balance
}

func (b *classicBank) Balance(account uuid.UUID) int64 {
	a := b.getAccount(account)
	return a.balance
}

func (b *classicBank) Transfer(from, to uuid.UUID, amount int64) {
	fromAccount := b.getAccount(from)
	toAccount := b.getAccount(to)
	fromAccount.balance = fromAccount.balance - amount
	toAccount.balance = toAccount.balance + amount
}

func (b *classicBank) Close() error {
	return nil
}
