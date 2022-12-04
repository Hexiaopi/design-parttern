package composite

import (
	"fmt"
)

type Component interface {
	execute()
}

type Leaf struct{}

func (l *Leaf) execute() {
	fmt.Println("leaf execute")
}

type Composite struct {
	children []Component
}

func (composite *Composite) execute() {
	for _, v := range composite.children {
		v.execute()
	}
}

func (composite *Composite) Add(c Component) {
	composite.children = append(composite.children, c)
}

func (composite *Composite) Remove(c Component) {
	flag := false
	switch c.(type) {
	case *Leaf:
		children := composite.children
		for i, v := range children {
			if value, ok := v.(*Leaf); ok {
				if value == c {
					flag = true
					switch i {
					case 0:
						composite.children = children[1:]
					case len(children) - 1:
						composite.children = children[0 : len(children)-1]
					default:
						composite.children = append(children[:i], children[i+1:]...)
					}
					break
				}
			}
		}
	case *Composite:
		children := composite.children
		for i, v := range children {
			if value, ok := v.(*Composite); ok {
				if value == c {
					flag = true
					switch i {
					case 0:
						composite.children = children[1:]
					case len(children) - 1:
						composite.children = children[0 : len(children)-1]
					default:
						composite.children = append(children[:i], children[i+1:]...)
					}
				}
			}
		}
	}
	if !flag {
		panic("not found")
	}
}

func (composite *Composite) GetChildren() []Component {
	return composite.children
}
