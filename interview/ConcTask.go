//Task on how simple mutex work

package main

import(
    "fmt"
	"sync"
	"sync/atomic"
)

type myMutex struct{
	locked int64
}

func (m *myMutex) Lock(){
	for {
		if atomic.CompareAndSwapInt64(&m.locked, 0, 1){
			return
		}
	}
}

func(m * myMutex) UnLock(){
	atomic.StoreInt64(&m.locked, 0)
}

func main(){
	wg := &sync.WaitGroup{}
	//mu:=&sync.Mutex{}
	mu:=myMutex{}
	c:= 0

	wg.Add(1000)

	for range 1000{
		go func(){
			defer wg.Done()
			mu.Lock()
			c++
			mu.UnLock()
		}()
	}

	wg.Wait()
	fmt.Println(c)
}
