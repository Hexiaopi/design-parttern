package builder

func ExampleDirector() {
	builder := ConcreteBuilder1{product: &Product{}}
	var director Director
	director.Director(builder)
	product := director.construct()
	product.show()
	// Output: A:A1	B:B1	C:C1
}


func ExampleDirector_second() {
	builder := ConcreteBuilder2{product: &Product{}}
	var director Director
	director.Director(builder)
	product := director.construct()
	product.show()
	// Output: A:A2	B:B2	C:C2
}
