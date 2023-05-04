package state

func ExampleState() {
	start := &StartState{}
	context := Context{state: start}
	context.Start()
	context.Run()
	context.Stop()
	context.Run()
	// Output:
	// already start
	// change start state to run
	// change run state to stop
	// can't change stop state to run
}
