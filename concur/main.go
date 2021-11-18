package main

import (
	"fmt"
	"time"
)

func doSomething(c chan string) {
	defer fmt.Println("Done with something")
	c <- "bob"
	c <- "frode"

}


func main() {
	fmt.Println("Hello there.")
	chan1 := make(chan string, 0)
	chan2 := make(chan string, 0)
	go doSomething(chan1)
	go doSomething(chan2)
	time.Sleep(time.Second)
	for {
		select {
			case m := <- chan1
			case m := <- chan1
			case <- time.After(10 * time.Second):
				fmt.Println("timeout")
		}
	}

	ret := <- someChan
	fmt.Println(ret)
	ret = <- someChan
	fmt.Println(ret)

}
