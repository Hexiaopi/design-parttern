package proxy

import "fmt"

type Subject interface {
	Request()
}

type RealSubject struct {
}

func (r RealSubject) Request() {
	fmt.Println("do request")
}

type Proxy struct {
	realSubject RealSubject
}

func (p Proxy) preRequest() {
	fmt.Println("pre request")
}

func (p Proxy) postRequest() {
	fmt.Println("post request")
}

func (p Proxy) Request() {
	p.preRequest()
	p.realSubject.Request()
	p.postRequest()
}
