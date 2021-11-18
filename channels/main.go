package main

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/perbu/go-intro/channels/bank"
	"github.com/perbu/go-intro/channels/channel"
	"github.com/perbu/go-intro/channels/classic"
	"github.com/perbu/go-intro/channels/naive"
	"log"
	"sync"
	"time"
)

const flavour = "channel"
const noOfAccounts = 5
const workers = 75
const name = "My Little Bank"

func main() {
	var myBank bank.Bank
	fmt.Printf("%s [%s]\n", name, flavour)

	switch flavour {
	case "classic":
		myBank = classic.CreateBank(name)
	case "channel":
		myBank = channel.CreateBank(name)
	case "naive":
		myBank = naive.CreateBank(name)
	default:
		panic("unknown bank flavour")
	}
	accounts := make([]uuid.UUID, 10)
	for i := 0; i < noOfAccounts; i++ {
		a := myBank.CreateAccount()
		accounts[i] = a
	}
	start := time.Now()
	wg := &sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		a := accounts[i % noOfAccounts]
		wg.Add(1)
		go runBench(myBank, a, wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Test took %s\n", elapsed)
	for i := 0; i < noOfAccounts; i++ {
		aid := accounts[i]
		balance := myBank.Balance(aid)
		if balance != 0 {
			fmt.Printf("Balance in acccount(%s) is wrong: %d\n", aid.String(), balance )
		}
	}
	err := myBank.Close()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(10*time.Millisecond)


}

func runBench(b bank.Bank, a uuid.UUID, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		b.Balance(a)
		b.Deposit(a, 100)
		for j := 0; j < 10; j++ {
			b.Withdraw(a, 10)
		}
	}
	wg.Done()
}
