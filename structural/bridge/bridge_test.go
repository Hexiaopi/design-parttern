package bridge

func ExampleOperation() {
	a := ConcreateImplA{}
	b := ConcreateImplB{}
	abstractA := Abstraction{Impl: a}
	abstractA.Operation()
	abstractB := Abstraction{Impl: b}
	abstractB.Operation()
	refineA := RefineAbstraction{Impl: a}
	refineA.Operation()
	refineB := RefineAbstraction{Impl: b}
	refineB.Operation()
	// Output:
	// Abstraction
	// ConcreateImplA
	// Abstraction
	// ConcreateImplB
	// RefineAbstraction
	// ConcreateImplA
	// RefineAbstraction
	// ConcreateImplB
}
