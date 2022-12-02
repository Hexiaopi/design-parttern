package facade

func ExampleOperation() {
	f := Facade{
		system1: Subsystem1{},
		system2: Subsystem2{},
		system3: Subsystem3{},
	}
	f.Operation()
	// Output:
	// subsystem1 operation
	// subsystem2 operation
	// subsystem3 operation
}
