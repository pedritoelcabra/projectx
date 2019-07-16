package core

import "github.com/hajimehoshi/ebiten"

type listenerFunction func(g *game)

type Input struct {
	listeners      map[string]map[string]listenerFunction
	rightMouseDown bool
	leftMouseDown  bool
}

func NewInput() *Input {
	aInput := &Input{}
	aInput.Init()
	return aInput
}

func (i *Input) Init() {
	i.listeners = make(map[string]map[string]listenerFunction)
	i.listeners["rightClick"] = make(map[string]listenerFunction)
	i.listeners["leftClick"] = make(map[string]listenerFunction)
}

func (i *Input) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		i.rightMouseDown = true
	} else {
		if i.rightMouseDown {
			for _, callback := range i.listeners["rightClick"] {
				callback(projectX)
			}
		}
		i.rightMouseDown = false
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		i.leftMouseDown = true
	} else {
		if i.leftMouseDown {
			for _, callback := range i.listeners["leftClick"] {
				callback(projectX)
			}
		}
		i.leftMouseDown = false
	}
}

func (i *Input) AddListener(event, name string, callback listenerFunction) {
	i.listeners[event][name] = callback
}
