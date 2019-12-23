package core

import "github.com/hajimehoshi/ebiten"

type listenerFunction func(g *game)

type Input struct {
	listeners      map[string]map[string]listenerFunction
	pressedKeys    map[ebiten.Key]bool
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
	i.listeners["RightClick"] = make(map[string]listenerFunction)
	i.listeners["LeftClick"] = make(map[string]listenerFunction)
	i.listeners["EscapePress"] = make(map[string]listenerFunction)
	i.listeners["EscapeRelease"] = make(map[string]listenerFunction)
	i.listeners["LeftPress"] = make(map[string]listenerFunction)
	i.listeners["LeftRelease"] = make(map[string]listenerFunction)
	i.listeners["DownPress"] = make(map[string]listenerFunction)
	i.listeners["DownRelease"] = make(map[string]listenerFunction)
	i.listeners["UpPress"] = make(map[string]listenerFunction)
	i.listeners["UpRelease"] = make(map[string]listenerFunction)
	i.listeners["RightPress"] = make(map[string]listenerFunction)
	i.listeners["RightRelease"] = make(map[string]listenerFunction)
	i.listeners["APress"] = make(map[string]listenerFunction)
	i.listeners["ARelease"] = make(map[string]listenerFunction)
	i.listeners["SPress"] = make(map[string]listenerFunction)
	i.listeners["SRelease"] = make(map[string]listenerFunction)
	i.listeners["WPress"] = make(map[string]listenerFunction)
	i.listeners["WRelease"] = make(map[string]listenerFunction)
	i.listeners["DPress"] = make(map[string]listenerFunction)
	i.listeners["DRelease"] = make(map[string]listenerFunction)
	i.pressedKeys = make(map[ebiten.Key]bool)
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		i.pressedKeys[k] = false
	}
}

func (i *Input) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		i.rightMouseDown = true
	} else {
		if i.rightMouseDown {
			i.TriggerCallbacks("RightClick")
		}
		i.rightMouseDown = false
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		i.leftMouseDown = true
	} else {
		if i.leftMouseDown {
			i.TriggerCallbacks("LeftClick")
		}
		i.leftMouseDown = false
	}
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			if !i.pressedKeys[k] {
				i.TriggerCallbacks(k.String() + "Press")
				i.pressedKeys[k] = true
			}
		} else {
			if i.pressedKeys[k] {
				i.TriggerCallbacks(k.String() + "Release")
				i.pressedKeys[k] = false
			}
		}
	}
}

func (i *Input) AddListener(event, name string, callback listenerFunction) {
	i.listeners[event][name] = callback
}

func (i *Input) TriggerCallbacks(key string) {
	for _, callback := range i.listeners[key] {
		callback(projectX)
	}
}
