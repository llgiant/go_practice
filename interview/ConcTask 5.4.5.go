package main

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

func unpredictableFunc() int {
	n := rand.Intn(40)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}

func predictableFunc(ctx context.Context) (int, error) {
	ch := make(chan struct{})
	var result int

	var cancel context.CancelFunc

	if _, ok := ctx.Deadline(); ok {
		ctx, cancel = context.WithDeadline(ctx, time.Now().Add(5*time.Second))
		defer cancel()
	}

	go func() {
		result = unpredictableFunc()
		close(ch)
	}()

	select {
	case <-ch:
		return result, nil
	case <-ctx.Done():
		return 0, errors.New("timeout exceeded")
	}
}

func main() {
	_, _ = predictableFunc(context.Background())
}
