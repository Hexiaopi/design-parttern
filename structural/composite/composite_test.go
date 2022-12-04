package composite

func ExampleLeaf() {
	var c Component
	l := Leaf{}
	c = &l
	c.execute()
	// Output:
	// leaf execute
}

func ExampleComposite() {
	var c Component
	composite := &Composite{
		children: make([]Component, 0),
	}
	composite.Add(&Leaf{})
	composite.Add(&Leaf{})
	c = composite
	c.execute()
	// Output:
	// leaf execute
	// leaf execute
}
