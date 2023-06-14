package bridge

func ExampleOperation() {
	a := ConcreateImplA{}
	b := ConcreateImplB{}
	var abstract Abstraction
	abstract = RefineAbstractionA{Impl: a}
	abstract.Operation()
	abstract = RefineAbstractionA{Impl: b}
	abstract.Operation()
	abstract = RefineAbstractionB{Impl: a}
	abstract.Operation()
	abstract = RefineAbstractionB{Impl: b}
	abstract.Operation()
	// Output:
	// RefineAbstractionA
	// ConcreateImplA
	// RefineAbstractionA
	// ConcreateImplB
	// RefineAbstractionB
	// ConcreateImplA
	// RefineAbstractionB
	// ConcreateImplB
}
