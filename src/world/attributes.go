package world

import (
	"github.com/pedritoelcabra/projectx/src/world/container"
	"log"
)

const (
	Vision AttributeKey = iota
	Speed
)

type AttributeKey int

type AttributeMap map[string]int

type Attributes struct {
	Data container.Container
}

func NewAttributes(initialValues AttributeMap) *Attributes {
	anAttributes := &Attributes{}
	anAttributes.SetValues(initialValues)
	return anAttributes
}

func (a *Attributes) Set(key AttributeKey, value int) {
	a.Data.Set(int(key), value)
}

func (a *Attributes) Get(key AttributeKey) int {
	return a.Data.Get(int(key))
}

func (a *Attributes) SetValues(values AttributeMap) {
	for key, value := range values {
		a.Set(GetAttributeKey(key), value)
	}
}

func GetAttributeKey(name string) AttributeKey {
	switch name {
	case "Vision":
		return Vision
	case "Speed":
		return Speed
	}
	log.Fatal("Invalid attribute: " + name)
	return 0
}
