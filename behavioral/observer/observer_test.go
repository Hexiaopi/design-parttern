package observer

func ExampleObserver() {
	ctx := Context{"test"}
	co1 := ConcreteObserver1{}
	co2 := ConcreteObserver2{}
	subject := Subject{
		ObserverCollection: make([]Observer, 0),
	}
	subject.Register(co1)
	subject.Register(co2)
	subject.NotifyObservers(ctx)
	ctx.Id = "other"
	subject.UnRegister(co1)
	subject.NotifyObservers(ctx)
	// Output:
	// concrete observer 1 update ctx: {test}
	// concrete observer 2 update ctx: {test}
	// concrete observer 2 update ctx: {other}
}
