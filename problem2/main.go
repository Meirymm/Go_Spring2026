package main
import (
	"fmt"
	"sync"
	"sync/atomic"
)
func atomicCounter() int {
	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	return int(counter)
}

func mutexCounter() int {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return counter
}

func main() {
	fmt.Println("=== Way 1: sync/atomic ===")
	fmt.Println("Result:", atomicCounter())
	fmt.Println("=== Way 2: sync.Mutex ===")
	fmt.Println("Result:", mutexCounter())
}