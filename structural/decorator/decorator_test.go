package decorator

func ExampleComponent() {
	c := ConcreteComponent{}
	c.execute()
	var c1 ConcreteDecortor1
	c1.Decorator.Decorator(c)
	c1.execute()
	var c2 ConcreteDecortor2
	c2.Decorator.Decorator(c)
	c2.execute()
	// Output:
	// concrete component execute
	// concrete decortor1 add function
	// concrete component execute
	// concrete decortor2 add function
	// concrete component execute
}
