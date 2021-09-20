package classic

import (
	uuid "github.com/nu7hatch/gouuid"
	"github.com/perbu/go-intro/channels/bank"
	"sync"
)

type account struct {
	balance float64
	mu      sync.RWMutex
}

type accountMap map[uuid.UUID]*account

type classicBank struct {
	name     string
	accounts accountMap
	mu       sync.RWMutex
}

func CreateBank(name string) bank.Bank {
	c := classicBank{
		name:     name,
		accounts: make(accountMap),
		mu:       sync.RWMutex{},
	}
	return &c
}

func (b *classicBank) CreateAccount() uuid.UUID {
	b.mu.Lock()
	defer b.mu.Unlock()
	accId, _ := uuid.NewV4()
	acc := &account{
		balance: 0.0,
	}
	b.accounts[*accId] = acc
	return *accId
}

func (b *classicBank) getAccount(account uuid.UUID) *account {
	b.mu.Lock()
	defer b.mu.Unlock()
	acct, _ := b.accounts[account]
	return acct
}

func (b *classicBank) Deposit(account uuid.UUID, amount float64) float64 {
	a := b.getAccount(account)
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance = a.balance + amount
	return a.balance
}

func (b *classicBank) Withdraw(account uuid.UUID, amount float64) float64 {
	a := b.getAccount(account)
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance = a.balance - amount
	return a.balance
}

func (b *classicBank) Balance(account uuid.UUID) float64 {
	a := b.getAccount(account)
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func (b *classicBank) Transfer(from, to uuid.UUID, amount float64) {
	fromAccount := b.getAccount(from)
	toAccount := b.getAccount(to)
	fromAccount.mu.Lock()
	toAccount.mu.Lock()
	fromAccount.balance = fromAccount.balance - amount
	toAccount.balance = toAccount.balance + amount
	toAccount.mu.Unlock()
	fromAccount.mu.Unlock()
}

func (b *classicBank) Close() error {
	return nil
}
