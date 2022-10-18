package adapter

func ExampleRequest() {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	target.Request()
	// Output:
	// adapteeImpl SpecificRequest
}
