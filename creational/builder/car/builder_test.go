package car

import "fmt"

func ExampleBuilder() {
	director := &Director{}
	benzBuilder := NewBenzBuilder()
	director.SetBuilder(benzBuilder)
	car := director.Construct()
	fmt.Println("Wheel:", car.Wheel)
	fmt.Println("Engine:", car.Engine)
	fmt.Println("Body:", car.Body)
	fmt.Println("Steering Wheel:", car.SteeringWheel)
	// Output:
	// Wheel: Benz Wheel
	// Engine: Benz Engine
	// Body: Benz Body
	// Steering Wheel: Benz Steering Wheel
	// Wheel: BMW Wheel
	// Engine: BMW Engine
	// Body: BMW Body
	// Steering Wheel: BMW Steering Wheel

	bmwBuilder := NewBMWBuilder()
	director.SetBuilder(bmwBuilder)
	car = director.Construct()
	fmt.Println("Wheel:", car.Wheel)
	fmt.Println("Engine:", car.Engine)
	fmt.Println("Body:", car.Body)
	fmt.Println("Steering Wheel:", car.SteeringWheel)
}
