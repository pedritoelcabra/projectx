package world

import (
	"github.com/pedritoelcabra/projectx/src/world/container"
	"log"
)

type AttributeMap map[string]int

type Attributes struct {
	Data *container.Container
}

func NewAttributes(initialValues AttributeMap) *Attributes {
	anAttributes := &Attributes{}
	anAttributes.Data = container.NewContainer()
	anAttributes.SetValues(initialValues)
	return anAttributes
}

func (a *Attributes) Set(key int, value int) {
	a.Data.Set(int(key), value)
}

func (a *Attributes) Apply(key int, value int) {
	a.Data.Set(int(key), value+a.Data.Get(int(key)))
}

func (a *Attributes) Get(key int) int {
	return a.Data.Get(int(key))
}

func (a *Attributes) SetValues(values AttributeMap) {
	for key, value := range values {
		a.Set(GetAttributeKey(key), value)
	}
}

func (a *Attributes) ApplyValues(values AttributeMap) {
	for key, value := range values {
		a.Apply(GetAttributeKey(key), value)
	}
}

func GetAttributeKey(name string) int {
	switch name {
	case "Vision":
		return Vision
	case "Speed":
		return Speed
	}
	log.Fatal("Invalid attribute: " + name)
	return 0
}
