package world

type World struct {
}

func NewWorld() *World {
	aWorld := &World{}
	aWorld.Init()
	return aWorld
}

func (w *World) Init() {

}
