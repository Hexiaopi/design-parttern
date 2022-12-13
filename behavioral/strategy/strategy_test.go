package strategy

func ExampleStrategy() {
	c := Context{}
	s1 := ConcreteStrategy1{}
	c.SetStrategy(s1)
	c.Do()
	s2 := ConcreteStrategy2{}
	c.SetStrategy(s2)
	c.Do()
	// Output:
	// do something
	// concrete strategy1 execute
	// do something
	// concrete strategy2 execute
}
