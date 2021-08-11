package proxy

func ExampleRequest() {
	proxy := Proxy{realSubject: RealSubject{}}
	proxy.Request()
	// Output:
	// pre request
	// do request
	// post request
}
