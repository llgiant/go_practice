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
	SendRequest_2(ctx context.Context, request Request) error
	WithLimiter(ctx context.Context, requests []Request)
}

type client struct {
}

func (c *client) SendRequest_2(ctx context.Context, request Request) error {
	// имитация отправки запроса
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Request sent with payload: " + request.Payload)
	return nil
}

// ограничение количества запросов в секунду RPS через берст
var rps = 1
var burst = 10

func (c *client) WithLimiter(ctx context.Context, reqs []Request) {
	ticker := time.NewTicker(time.Second / time.Duration(rps))
	tickets := make(chan Request, burst)

	wg := sync.WaitGroup{}

	go func() {
		for {
			select {
			case <-ticker.C:
				tickets <- struct{}{}
			}
		}
	}()

	wg.Add(len(reqs))
	for _, req := range reqs {
		<-ticker.C
		go func() {
			defer wg.Done()
			c.SendRequest_2(ctx, req)
		}()
	}
	wg.Wait()
}

func main() {
	ctx := context.Background()
	c := &client{}
	reqs := make([]Request, 1000)
	for i := 0; i < 1000; i++ {
		reqs[i] = Request{Payload: "Request #" + strconv.Itoa(i+1)}
	}
	c.WithLimiter(ctx, reqs)
}
