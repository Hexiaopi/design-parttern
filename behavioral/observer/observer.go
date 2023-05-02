package observer

import "fmt"

type Context struct {
	Id string
}

type Observer interface {
	Update(ctx Context)
}

type ConcreteObserver1 struct {
}

func (c ConcreteObserver1) Update(ctx Context) {
	fmt.Println("concrete observer 1 update ctx:", ctx)
}

type ConcreteObserver2 struct {
}

func (c ConcreteObserver2) Update(ctx Context) {
	fmt.Println("concrete observer 2 update ctx:", ctx)
}

type Subject struct {
	ObserverCollection []Observer
}

func (s *Subject) Register(observer Observer) {
	s.ObserverCollection = append(s.ObserverCollection, observer)
}

func (s *Subject) UnRegister(observer Observer) {
	for i := 0; i < len(s.ObserverCollection); i++ {
		if s.ObserverCollection[i] == observer {
			s.ObserverCollection = append(s.ObserverCollection[:i], s.ObserverCollection[i+1:]...)
			break
		}
	}
}

func (s *Subject) NotifyObservers(ctx Context) {
	for i := 0; i < len(s.ObserverCollection); i++ {
		s.ObserverCollection[i].Update(ctx)
	}
}
