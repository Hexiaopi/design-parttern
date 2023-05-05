package visitor

import "fmt"

func ExampleShape() {
	square := &Square{SideLength: 4.0}
	rectangle := &Rectangle{Width: 3.0, Height: 5.0}
	circle := &Circle{Radius: 2.0}

	areaVisitor := &AreaVisitor{}
	square.Accept(areaVisitor)
	fmt.Println("square area:", areaVisitor.Area)
	rectangle.Accept(areaVisitor)
	fmt.Println("rectangle area:", areaVisitor.Area)
	circle.Accept(areaVisitor)
	fmt.Println("circle area:", areaVisitor.Area)

	perimeterVisitor := &PerimeterVisitor{}
	square.Accept(perimeterVisitor)
	fmt.Println("square perimeter:", perimeterVisitor.Perimeter)
	rectangle.Accept(perimeterVisitor)
	fmt.Println("rectangle perimeter:", perimeterVisitor.Perimeter)
	circle.Accept(perimeterVisitor)
	fmt.Println("circle perimeter:", perimeterVisitor.Perimeter)

	// Output:
	// square area: 16
	// rectangle area: 15
	// circle area: 12.566370614359172
	// square perimeter: 16
	// rectangle perimeter: 16
	// circle perimeter: 12.566370614359172
}
