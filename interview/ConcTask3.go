package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main(){
	//counter:=0
	var counter int64 = 0
	wg:=sync.WaitGroup{}
	//mu:=sync.Mutex{}


	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			//mu.Lock()
			//defer mu.Unlock()

			//counter+=1
			atomic.AddInt64(&counter,1)
		}()
	}

	wg.Wait()
		fmt.Println(counter)

}