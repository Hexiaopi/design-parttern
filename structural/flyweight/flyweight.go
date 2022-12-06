package flyweight

import "fmt"

type UnsharedFlyweight struct {
	ExternalState string
}

type ConcreteFlyweight struct {
	InternalState string
}

func (concrete ConcreteFlyweight) execute(unshared UnsharedFlyweight) {
	fmt.Println(concrete.InternalState)
	fmt.Println(unshared.ExternalState)
}

type Flyweight interface {
	execute(unshared UnsharedFlyweight)
}

type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func (factory FlyweightFactory) GetFlyweight(key string) Flyweight {
	if flyweight, ok := factory.flyweights[key]; !ok {
		return nil
	} else {
		return flyweight
	}
}
