package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/world"
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
		ProjectX.HandleAttackClick()
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
	i.InitKey(event)
	i.listeners[event][name] = callback
}

func (i *Input) InitKey(event string) {
	if _, ok := i.listeners[event]; !ok {
		i.listeners[event] = make(map[string]listenerFunction)
	}
}

func (i *Input) TriggerCallbacks(key string) {
	for _, callback := range i.listeners[key] {
		callback(ProjectX)
	}
}

func (g *game) InitInput() {
	g.Input = NewInput()
	g.Input.AddListener("RightClick", "rightClick", func(g *game) {
		if g.selectedBuilding != nil {
			g.selectedBuilding = nil
			g.Gui.SetDisabled(EntityMenu, true)
		}
		g.Gui.AddMenu(ContextMenu, g.BuildContextMenu(ebiten.CursorPosition()))
	})
	g.Input.AddListener("LeftClick", "leftClick", func(g *game) {
		g.Gui.SetDisabled(ContextMenu, true)
		g.OpenEntityMenu(g.MousePosCoord().XY())
	})
	g.Input.AddListener("EscapePress", "toggleMenu", func(g *game) {
		if g.Gui.GetMenu(ContextMenu) != nil && !g.Gui.GetMenu("context").IsDisabled() {
			g.Gui.SetDisabled(ContextMenu, true)
			return
		}
		if !g.Gui.GetMenu(BuildingMenu).IsDisabled() {
			g.Gui.SetDisabled(BuildingMenu, true)
			return
		}
		if !g.Gui.GetMenu(EntityMenu).IsDisabled() {
			g.Gui.SetDisabled(EntityMenu, true)
			g.selectedBuilding = nil
			return
		}
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
	g.Input.AddListener("GraveAccentPress", "toggleDebug", func(g *game) {
		g.Gui.ToggleDebug()
	})
	g.Input.AddListener("F5Press", "quickSave", func(g *game) {
		g.QuickSave()
	})
	g.Input.AddListener("F9Press", "quickLoad", func(g *game) {
		g.QuickLoad()
	})
	g.Input.AddListener("F1Press", "toggleDebug", func(g *game) {
		g.Gui.ToggleDebug()
	})
}
