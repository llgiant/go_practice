package main

import (
	"strconv"
"time"
"fmt"
"context")

type Request struct {
	Payload string
}

type Client interface{
	SendRequest( ctx context.Context,request Request) error
	WithLimiter(ctx context.Context, requests []Request) 
}

type client struct {

}

func (c *client) SendRequest( ctx context.Context,request Request) error{
	// имитация отправки запроса
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Request sent with payload: " + request.Payload)
	return nil
}

//ограничение количества запросов в секунду
func (c *client) WithLimiter(ctx context.Context, requests []Request) {}

func main() {
	ctx := context.Background()
	c := &client{}
	requests := make([]Request,1000)
	for i := 0; i < 1000; i++ {
		requests[i] = Request{Payload: "Request #" + strconv.Itoa(i+1)}
	}
	c.WithLimiter(ctx, requests)
}