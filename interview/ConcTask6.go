package main

import (
	"fmt"

)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {	
			ch <- i
		}
		close(ch) // close the channel after sending all values
		//otherwise, the range loop below will block forever and cause a deadlock
	}()
	for val := range ch {
		fmt.Println(val)
	}
}