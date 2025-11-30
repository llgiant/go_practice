package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type Request struct {
	Payload string
}

type Client interface {
	SendRequest(ctx context.Context, request Request) error
	WithLimiter(ctx context.Context, requests []Request)
}

type client struct {
}

func (c *client) SendRequest(ctx context.Context, request Request) error {
	// имитация отправки запроса
	time.Sleep(1 * time.Second)
	fmt.Println("Request sent with payload: " + request.Payload)
	return nil
}

// ограничение работающих горутин
var maxGoroutines = 100

func (c *client) WithLimiter(ctx context.Context, requests []Request) {
		tokens := make(chan struct{}, maxGoroutines)

	go func ()  {
		for range maxGoroutines {
			tokens <- struct{}{}
		}
	}()

	for _, req := range requests {
		<-tokens

		go func() {
			defer func() { tokens <- struct{}{} }()
			c.SendRequest(ctx, req)
		}()
		}


		for range maxGoroutines {
			<-tokens
		}
}

func main() {
	ctx := context.Background()
	c := &client{}
	requests := make([]Request, 1000)
	for i := 0; i < 1000; i++ {
		requests[i] = Request{Payload: "Request #" + strconv.Itoa(i+1)}
	}
	c.WithLimiter(ctx, requests)
}
