package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	count := 0

	goroutineMax := 10
	wg.Add(goroutineMax)
	for i := 0; i < goroutineMax; i++ {
		go func() {
			mu.Lock()
			temp := count
			runtime.Gosched()
			temp++
			count = temp
			fmt.Println(count)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter >", count)

}
