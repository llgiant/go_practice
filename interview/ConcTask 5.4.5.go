package main

import (
	"crypto/rand"
	"errors"
	"time"
	"math/rand"
)

  func unpredictableFunc() int {
	n:=rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}

func predictableFunc(ctx context.Context) (int, error) {
ch:=make(chan struct{})

go func (){
	result:=unpredictableFunc()
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
_,_ = predictableFunc(context.Background())	
  }