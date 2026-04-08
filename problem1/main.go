package main
import (
	"fmt"
	"sync"
)
func withSyncMap() {
	var sm sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Store("key", i)
		}(i)
	}
	wg.Wait()
	value, _ := sm.Load("key")
	fmt.Printf("sync.Map result: %v\n", value)
}

func withRWMutex() {
	m := make(map[string]int)
	var mu sync.RWMutex
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m["key"] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	mu.RLock()
	value := m["key"]
	mu.RUnlock()
	fmt.Printf("RWMutex result: %v\n", value)
}

func main() {
	fmt.Println("=== Way 1: sync.Map ===")
	withSyncMap()
	fmt.Println("=== Way 2: sync.RWMutex ===")
	withRWMutex()
}