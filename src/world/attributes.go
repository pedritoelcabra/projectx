package world

import (
	"github.com/pedritoelcabra/projectx/src/world/container"
	"log"
)

type AttributeMap map[string]int

type Attributes struct {
	Data *container.Container
}

func NewEmptyAttributes() *Attributes {
	emptyMap := make(AttributeMap)
	return NewAttributes(emptyMap)
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

func (a *Attributes) SetF(key int, value float64) {
	a.Data.SetF(key, value)
}

func (a *Attributes) ApplyF(key int, value float64) {
	a.Data.SetF(int(key), value+a.Data.GetF(int(key)))
}

func (a *Attributes) Get(key int) int {
	return a.Data.Get(int(key))
}

func (a *Attributes) GetF(key int) float64 {
	return a.Data.GetF(key)
}

func (a *Attributes) SetValues(values AttributeMap) {
	for key, value := range values {
		a.SetF(GetAttributeKey(key), float64(value))
	}
}

func (a *Attributes) ApplyValuesF(values AttributeMap) {
	for key, value := range values {
		a.ApplyF(GetAttributeKey(key), float64(value))
	}
}

func GetAttributeKey(name string) int {
	switch name {
	case "Vision":
		return Vision
	case "Speed":
		return Speed
	case "Size":
		return Size
	case "AttackRange":
		return AttackRange
	case "AttackDamage":
		return AttackDamage
	case "AttackSpeed":
		return AttackSpeed
	case "MaxHitPoints":
		return MaxHitPoints
	case "SpawnsWild":
		return SpawnsWild
	}
	log.Fatal("Invalid attribute: " + name)
	return 0
}
