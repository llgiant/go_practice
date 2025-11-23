package main

import (
	"fmt"
)

func worker(n int) chan string{
		ch := make(chan string,1)
		go func() { //use goroutine to send messages 
			for i := 0; i < n; i++ {
		ch <- fmt.Sprintf("msg %d", i)
	}

	close(ch) //close the channel after sending all values eventhough it's buffered
		}()
	
	return ch
}
func main() {

n:=10
for msg := range worker(n) {
		fmt.Println("revieved:", msg)
	}
}