package main
import (
	"fmt"
	"time"
)

func worker() <-chan int {
	result := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		close(result)
	}()

	return result
}

func main() {
	timeStart := time.Now()
	_,_=<-worker(),<-worker()
	fmt.Println(time.Since(timeStart))
	newtimeStart := time.Now()
	_,_= worker(), worker()
	fmt.Println(time.Since(newtimeStart))
}