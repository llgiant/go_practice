package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
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
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Request sent with payload: " + request.Payload)
	return nil
}

// ограничение количестова коннктов
var maxConnects = 10

func (c client) WithLimiter(ctx context.Context, ch chan Request) {
	wg := sync.WaitGroup{}
	wg.Add(maxConnects)
	for range maxConnects {
		go func() {
			defer wg.Done()
			for req := range ch {
				c.SendRequest(ctx, req)
			}
		}()
	}
	wg.Wait()
}

func main() {
	ctx := context.Background()
	c := &client{}
	requests := make([]Request, 1000)
	for i := 0; i < 1000; i++ {
		requests[i] = Request{Payload: "Request #" + strconv.Itoa(i+1)}
	}
	c.WithLimiter(ctx, generate(requests))
}

func generate(reqs []Request) chan Request {
	ch := make(chan Request)
	go func() {
		for _, req := range reqs {
			ch <- req
		}
		close(ch)
	}()
	return ch
}
