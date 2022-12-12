package iterator

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type ConcreteIterator struct {
	numbers []string
	next    int
}

func (i *ConcreteIterator) HasNext() bool {
	return i.next < len(i.numbers)
}

func (i *ConcreteIterator) Next() interface{} {
	if i.HasNext() {
		numer := i.numbers[i.next]
		i.next += 1
		return numer
	}
	return nil
}

type ConcreteAggregate struct {
	Length int
}

func (a *ConcreteAggregate) Iterator() Iterator {
	return &ConcreteIterator{
		numbers: make([]string, a.Length),
		next:    0,
	}
}
