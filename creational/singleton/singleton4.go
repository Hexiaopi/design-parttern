package singleton

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type single4 struct{}

var rwlock = &sync.Mutex{}
var singleInstance4 *single4
var flag uint32

func GetInstance4() *single4 {
	if atomic.LoadUint32(&flag) == 0 {
		rwlock.Lock()
		defer rwlock.Unlock()
		if atomic.LoadUint32(&flag) == 0 {
			fmt.Println("Creating Single Instance Now")
			singleInstance1 = &single1{}
			defer atomic.StoreUint32(&flag, 1)
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance4
}
