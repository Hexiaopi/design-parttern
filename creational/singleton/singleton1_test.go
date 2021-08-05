package singleton

import (
	"sync"
	"testing"
)

func TestGetInstance1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			GetInstance1()
		}()
	}
	wg.Wait()
}
