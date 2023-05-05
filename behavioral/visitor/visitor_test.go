package visitor

func ExampleVisitor() {
	visitor := &ConcreteVisitor{}
	elementA := &ConcreteElementA{}
	elementB := &ConcreteElementB{}
	elementA.Accept(visitor)
	elementB.Accept(visitor)
	// Output:
	// Visit ConcreteElementA
	// Visit ConcreteElementB
}
