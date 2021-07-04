package abstract_factory

func ExampleConcreteFactory1_NewProduct() {
	var factory AbstractFactory
	factory = ConcreteFactory1{}
	product1 := factory.NewProduct1()
	Product2 := factory.NewProduct2()
	product1.Show()
	Product2.Use()
	// Output:
	// Product11 show
	// Product12 use
}

func ExampleConcreteFactory2_NewProduct() {
	var factory AbstractFactory
	factory = ConcreteFactory2{}
	product1 := factory.NewProduct1()
	product2 := factory.NewProduct2()
	product1.Show()
	product2.Use()
	// Output:
	// Product21 show
	// Product22 use
}
