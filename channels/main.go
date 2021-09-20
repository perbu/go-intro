package main

import (
	uuid "github.com/nu7hatch/gouuid"
	"github.com/perbu/go-intro/channels/bank"
	"github.com/perbu/go-intro/channels/channel"
	"github.com/perbu/go-intro/channels/classic"
	"log"
	"sync"
	"time"
)

const useChannel = false

func main() {
	var myLittleBank bank.Bank
	const name = "My Little Bank"
	if useChannel {
		myLittleBank = channel.CreateBank(name)
	} else  {
		myLittleBank = classic.CreateBank(name)
	}
	accounts := make([]uuid.UUID,10)
	for i := 0; i<10; i++ {
		accounts[i] = myLittleBank.CreateAccount()
	}
	start := time.Now()
	wg := &sync.WaitGroup{}
	for i := 0; i<10; i++ {
		a:=accounts[i]
		wg.Add(1)
		go runBench(myLittleBank, a, wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Test took %s\n", elapsed)
}

func runBench(b bank.Bank, a uuid.UUID, wg *sync.WaitGroup) {
	for i := 0; i<10000; i++ {
		b.Balance(a)
		b.Deposit(a, 100)
		for j := 0; j<10; j++ {
			b.Withdraw(a, 10)
		}
	}
	wg.Done()
}



