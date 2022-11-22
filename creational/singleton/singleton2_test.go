package singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			instance2 := GetInstance2()
			fmt.Printf("get single instance address %p and is nil? %t\n", instance2, instance2 == nil)
		}()
	}
	wg.Wait()
}


func BenchmarkGetInstance2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetInstance2()
	}
}