package factory_method

func ExampleConcreteFactory1_NewProduct() {
	var factory AbstractFactory
	factory = ConcreteFactory1{}
	product := factory.NewProduct()
	product.Show()
	// Output:
	// Product1 show
}

func ExampleConcreteFactory2_NewProduct() {
	var factory AbstractFactory
	factory = ConcreteFactory2{}
	product := factory.NewProduct()
	product.Show()
	// Output:
	// Product2 show
}
