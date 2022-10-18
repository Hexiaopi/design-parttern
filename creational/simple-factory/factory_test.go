package simple_factory

func ExampleSimpleFactory_NewProduct1() {
	factory := SimpleFactory{}
	product := factory.NewProduct(Product1)
	product.Show()
	// Output:
	// Product1 show
}

func ExampleSimpleFactory_NewProduct2() {
	factory := SimpleFactory{}
	product := factory.NewProduct(Product2)
	product.Show()
	// Output:
	// Product2 show
}
