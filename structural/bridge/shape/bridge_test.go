package shape

func ExampleBridge() {
	red := Red{}
	blue := Blue{}

	var shape Shape
	shape = Circle{red}
	shape.Draw()
	shape = Circle{blue}
	shape.Draw()
	shape = Square{red}
	shape.Draw()
	shape = Square{blue}
	shape.Draw()
	// Output:
	// red circle
	// blue circle
	// red square
	// blue square
}
