package object_pool

import (
	"sync"
)

type pool struct {
	mutex     *sync.Mutex
	idle      []interface{}
	active    []interface{}
	newObject func() interface{}
}

func NewPool(newFunc func() interface{}) *pool {
	return &pool{
		mutex:     &sync.Mutex{},
		idle:      make([]interface{}, 0),
		active:    make([]interface{}, 0),
		newObject: newFunc,
	}
}

func (p *pool) Acquire() interface{} {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	var object interface{}
	if len(p.idle) == 0 {
		object = p.newObject()
	} else {
		object = p.idle[0]
		p.idle = p.idle[1:]
	}
	p.active = append(p.active, object)
	return object
}

func (p *pool) Release(target interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.idle = append(p.idle, target)
	for i := range p.active {
		if p.active[i] == target {
			p.active = append(p.active[:i], p.active[i+1:]...)
			break
		}
	}
}
