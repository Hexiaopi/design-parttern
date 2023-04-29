package mediator

func ExampleMediator() {
	mediator := ConcreteMediator{}
	colleague1 := ConcreteColleague1{mediator: mediator}
	colleague2 := ConcreteColleague2{mediator: mediator}
	mediator.ConcreteColleague1 = colleague1
	mediator.ConcreteColleague2 = colleague2

	colleague1.SendMessage("hello")
	colleague2.SendMessage("hi")
	// Output:
	// ConcreateColleague1 send message: hello
	// ConcreateColleague2 receive message: hello
	// ConcreateColleague2 send message: hi
	// ConcreateColleague1 receive message: hi
}
