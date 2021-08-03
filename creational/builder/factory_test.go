package builder

func ExampleDirector_Builder1() {
	builder := ConcreteBuilder1{product: &Product{}}
	var director Director
	director.Director(builder)
	product := director.construct()
	product.show()
	// Output:
	// A:A1	B:B1	C:C1
}


func ExampleDirector_Builder2() {
	builder := ConcreteBuilder2{product: &Product{}}
	var director Director
	director.Director(builder)
	product := director.construct()
	product.show()
	// Output:
	// A:A2	B:B2	C:C2
}
