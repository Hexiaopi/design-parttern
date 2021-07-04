package abstract_factory

func ExampleRDBFactory() {
	factory := &RDBDaoFactory{}
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
	// Output:
	// rdb main save
	// rdb detail save
}

func ExampleXMLDaoFactory() {
	factory := &XMLDaoFactory{}
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
	// Output:
	// xml main save
	// xml detail save
}
