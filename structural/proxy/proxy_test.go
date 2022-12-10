package proxy

func ExampleRequest() {
	var subject Subject
	subject = Proxy{realSubject: RealSubject{}}
	subject.Request()
	// Output:
	// pre request
	// do request
	// post request
}
