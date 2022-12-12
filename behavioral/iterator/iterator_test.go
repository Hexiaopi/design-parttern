package iterator

import "fmt"

func ExampleIterator() {
	var aggregate Aggregate
	aggregate = &ConcreteAggregate{Length: 10}
	iterator := aggregate.Iterator()
	i := 0
	for iterator.HasNext() {
		_ = iterator.Next()
		fmt.Println(i)
		i++
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
}
