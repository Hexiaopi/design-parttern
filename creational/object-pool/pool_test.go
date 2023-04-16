package object_pool

import (
	"fmt"
)

func ExamplePool() {
	newConnDB := func() interface{} {
		return "conn-db"
	}
	pool := NewPool(newConnDB)
	object1 := pool.Acquire()
	fmt.Println(len(pool.idle))
	fmt.Println(len(pool.active))

	object2 := pool.Acquire()
	fmt.Println(len(pool.idle))
	fmt.Println(len(pool.active))

	pool.Release(object1)
	fmt.Println(len(pool.idle))
	fmt.Println(len(pool.active))

	pool3 := pool.Acquire()
	fmt.Println(len(pool.idle))
	fmt.Println(len(pool.active))

	pool.Release(object2)
	pool.Release(pool3)
	fmt.Println(len(pool.idle))
	fmt.Println(len(pool.active))
	// Output:
	// 0
	// 1
	// 0
	// 2
	// 1
	// 1
	// 0
	// 2
	// 2
	// 0
}
