package bank

import uuid "github.com/nu7hatch/gouuid"

type Bank interface {
	CreateAccount() uuid.UUID
	Deposit(account uuid.UUID, amount int64) int64
	Withdraw(account uuid.UUID, amount int64) int64
	Balance(account uuid.UUID) int64
	Transfer(from, to uuid.UUID, amount int64)
	Close() error
}
