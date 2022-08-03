package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("Hello, world!")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	timeoutCtx, tmOutCancel := context.WithTimeout(ctx, 3*time.Second)
	defer tmOutCancel()
	run(timeoutCtx, 10*time.Second)
}

func run(ctx context.Context, sleepDuration time.Duration) {
	start := time.Now()
	err := sleepContext(ctx, sleepDuration)
	fmt.Printf("slept for %v (err = %v) \n", time.Since(start), err)
}

func sleepContext(ctx context.Context, dur time.Duration) error {
	// sleep for a while
	select {
	case <-time.After(dur):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
