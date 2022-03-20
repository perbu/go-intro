package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type stateMachine struct {
	a           int
	b           int
	submissions fnChan
}
type fnChan chan func()

func (s stateMachine) Run(ctx context.Context) {
loop:
	for {
		select {
		case <-ctx.Done():
			close(s.submissions)
			break loop
		case fn := <-s.submissions:
			fn()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	s := stateMachine{
		submissions: make(fnChan, 0),
	}
	wg.Add(1)
	go func() {
		s.Run(ctx)
		wg.Done()
	}()
	resCh := make(chan int)
	s.submissions <- func() {
		s.a = 3
		s.b = 4
	}
	s.submissions <- func() {
		resCh <- s.a + s.b
	}
	fmt.Println("First result: ", <-resCh)
	s.submissions <- func() {
		s.a = 5
		s.b = 5
	}
	s.submissions <- func() {
		resCh <- s.a + s.b
	}
	fmt.Println("Second res: ", <-resCh)

	cancel()
	wg.Wait()
}
