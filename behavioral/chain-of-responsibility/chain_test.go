package chainofresponsibility

func ExampleHandler() {
	h1 := ConcreteHandler1{}
	h2 := ConcreteHandler2{}
	h1.SetNext(&h2)
	h1.Handle()
	// Output:
	// concrete handler1
	// concrete handler2
}
