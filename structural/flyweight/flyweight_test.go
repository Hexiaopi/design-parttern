package flyweight

func ExampleFlyweight() {
	factory := FlyweightFactory{flyweights: map[string]Flyweight{
		"a": ConcreteFlyweight{InternalState: "a"},
		"b": ConcreteFlyweight{InternalState: "bb"},
		"c": ConcreteFlyweight{InternalState: "ccc"},
	}}
	a := factory.GetFlyweight("a")
	a.execute(UnsharedFlyweight{ExternalState: "1"})
	b := factory.GetFlyweight("b")
	b.execute(UnsharedFlyweight{ExternalState: "2"})
	c := factory.GetFlyweight("c")
	c.execute(UnsharedFlyweight{ExternalState: "3"})
	// Output:
	// a
	// 1
	// bb
	// 2
	// ccc
	// 3
}
