package main

import "sync"

func fanin(chans ...<-chan int) <-chan int {
	result := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		for _, ch := range chans {
			wg.Add(1)
			go func(ch <-chan int) {
				defer wg.Done()
				for val := range ch {
					result <- val
				}
			}(ch)
		}
		wg.Wait()
		close(result)
	}()
	return result
}
