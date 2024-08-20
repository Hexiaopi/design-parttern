package decorator

func ExampleComponent() {
	c := ConcreteComponent{}
	c.execute()
	//c1装饰c
	var c1 ConcreteDecortor1
	c1.SetComponent(c)
	c1.execute()
	//c2装饰c
	var c2 ConcreteDecortor2
	c2.SetComponent(c)
	c2.execute()
	//c2装饰c1
	var c21 ConcreteDecortor2
	c21.SetComponent(c1)
	c21.execute()
	// Output:
	// concrete component execute
	// concrete decortor1 add function
	// concrete component execute
	// concrete decortor2 add function
	// concrete component execute
	// concrete decortor2 add function
	// concrete decortor1 add function
	// concrete component execute
}
