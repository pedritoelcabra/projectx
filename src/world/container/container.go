package container

type Container struct {
	Values  map[int]int
	ValuesF map[int]float64
}

func NewContainer() *Container {
	aContainer := &Container{}
	aContainer.Values = make(map[int]int)
	aContainer.ValuesF = make(map[int]float64)
	return aContainer
}

func (c *Container) Get(key int) int {
	return c.Values[key]
}

func (c *Container) GetF(key int) float64 {
	return c.ValuesF[key]
}

func (c *Container) Set(key, value int) {
	c.Values[key] = value
}

func (c *Container) SetF(key int, value float64) {
	c.ValuesF[key] = value
}
