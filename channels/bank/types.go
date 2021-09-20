package bank

import uuid "github.com/nu7hatch/gouuid"

type Bank interface {
	CreateAccount() uuid.UUID
	Deposit(account uuid.UUID, amount float64) float64
	Withdraw(account uuid.UUID, amount float64) float64
	Balance(account uuid.UUID) float64
	Transfer(from, to uuid.UUID, amount float64)
	Close() error
}
