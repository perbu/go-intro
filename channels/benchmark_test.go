package main

import (
	"github.com/perbu/go-intro/channels/channel"
	"github.com/perbu/go-intro/channels/classic"
	"testing"
)

func BenchmarkChannel(b *testing.B) {
	bank := channel.CreateBank("Channel Bank Test")
	a := bank.CreateAccount()
	for n := 0; n < b.N; n++ {
		bank.Balance(a)
		bank.Deposit(a, 100)
		for j := 0; j<10; j++ {
			bank.Withdraw(a, 10)
		}
	}
}

func BenchmarkClassic(b *testing.B) {
	bank := classic.CreateBank("Channel Bank Test")
	a := bank.CreateAccount()
	for n := 0; n < b.N; n++ {
		bank.Balance(a)
		bank.Deposit(a, 100)
		for j := 0; j<10; j++ {
			bank.Withdraw(a, 10)
		}
	}
}


