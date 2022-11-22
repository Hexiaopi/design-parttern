package singleton

import (
	"sync"
	"testing"
)

func TestGetInstance4(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			GetInstance4()
		}()
	}
	wg.Wait()
}

func BenchmarkGetInstance4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetInstance4()
	}
}
