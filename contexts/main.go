package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	println("Hello, world!")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	run(ctx)
}

func run(ctx context.Context) {
	start := time.Now()
	err := sleepContext(ctx, 10*time.Second)
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
