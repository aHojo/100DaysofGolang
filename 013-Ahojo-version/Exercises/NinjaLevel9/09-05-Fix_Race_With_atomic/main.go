package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var count int64 = 0

	goroutineMax := 10
	wg.Add(goroutineMax)
	for i := 0; i < goroutineMax; i++ {
		go func() {

			atomic.AddInt64(&count, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&count))

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter >", count)

}
