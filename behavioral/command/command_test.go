package command

func ExampleCommand() {
	receiver := Receiver{}
	invoker := Invoker{}

	cmd1 := ConcreateCommand1{name: "admin", receiver: &receiver}
	invoker.SetCommand(&cmd1)
	invoker.ExecuteCommand()

	cmd2 := ConcreateCommand2{name: "admin", desc: "描述", address: "地址", receiver: &receiver}
	invoker.SetCommand(&cmd2)
	invoker.ExecuteCommand()

	// Output:
	// operation1: admin
	// operation2: admin 描述 地址
}