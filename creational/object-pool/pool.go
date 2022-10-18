package object_pool

import (
	"fmt"
	"sync"
)

type iPoolObject interface {
	getID() string
}

type pool struct {
	idle     []iPoolObject
	active   []iPoolObject
	capacity int
	mutex    *sync.Mutex
}

func initPool(poolObjects []iPoolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("Connot create a pool of 0 length")
	}
	active := make([]iPoolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mutex:    new(sync.Mutex),
	}
	return pool, nil
}

func (p *pool) loan() (iPoolObject, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No Pool Object Free")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	return obj, nil
}

func (p *pool) remove(target iPoolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("Target poll pbject doesn't belong to the pool")
}

func (p *pool) receive(target iPoolObject) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("Return Pool Object with ID: %s\n", target.getID())
	return nil
}

type connection struct {
	id string
}

func (c *connection) getID() string {
	return c.id
}
