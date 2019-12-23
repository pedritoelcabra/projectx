package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/world"
)

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

func (g *game) InitInput() {
	g.Input = NewInput()
	g.Input.AddListener("RightClick", "openContext", func(g *game) {
		g.Gui.AddMenu("context", g.BuildContextMenu(ebiten.CursorPosition()))
	})
	g.Input.AddListener("LeftClick", "closeContext", func(g *game) {
		g.Gui.SetDisabled("context", true)
	})
	g.Input.AddListener("EscapePress", "toggleMenu", func(g *game) {
		g.TogglePause()
	})
	g.Input.AddListener("LeftPress", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERLEFT, true)
	})
	g.Input.AddListener("LeftRelease", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERLEFT, false)
	})
	g.Input.AddListener("APress", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERLEFT, true)
	})
	g.Input.AddListener("ARelease", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERLEFT, false)
	})
	g.Input.AddListener("RightPress", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERRIGHT, true)
	})
	g.Input.AddListener("RightRelease", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERRIGHT, false)
	})
	g.Input.AddListener("DPress", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERRIGHT, true)
	})
	g.Input.AddListener("DRelease", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERRIGHT, false)
	})
	g.Input.AddListener("UpPress", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERUP, true)
	})
	g.Input.AddListener("UpRelease", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERUP, false)
	})
	g.Input.AddListener("WPress", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERUP, true)
	})
	g.Input.AddListener("WRelease", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERUP, false)
	})
	g.Input.AddListener("DownPress", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERDOWN, true)
	})
	g.Input.AddListener("DownRelease", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERDOWN, false)
	})
	g.Input.AddListener("SPress", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERDOWN, true)
	})
	g.Input.AddListener("SRelease", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERDOWN, false)
	})
}
