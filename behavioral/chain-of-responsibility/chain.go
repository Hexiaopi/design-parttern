package chainofresponsibility

import "fmt"

type Handler interface {
	Handle()
	SetNext(Handler)
}

type ConcreteHandler1 struct {
	next Handler
}

func (h *ConcreteHandler1) Handle() {
	fmt.Println("concrete handler1")
	if h.next != nil {
		h.next.Handle()
	}
}

func (h *ConcreteHandler1) SetNext(handler Handler) {
	h.next = handler
}

type ConcreteHandler2 struct {
	next Handler
}

func (h *ConcreteHandler2) Handle() {
	fmt.Println("concrete handler2")
	if h.next != nil {
		h.next.Handle()
	}
}

func (h *ConcreteHandler2) SetNext(handler Handler) {
	h.next = handler
}
