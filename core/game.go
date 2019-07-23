package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/gui"
	"github.com/pedritoelcabra/projectx/world"
	"strconv"
)

type game struct {
	Gui            *gui.Gui
	Input          *Input
	World          *world.World
	Graphics       *gfx.Graphics
	Screen         *gfx.Screen
	tick           int
	framesDrawn    int
	isPaused       bool
	rightMouseDown bool
	debugMessage   string
}

var projectX *game

func New() *game {
	aGame := game{}
	aGame.init()
	return &aGame
}

func G() *game {
	if projectX == nil {
		projectX = New()
	}
	return projectX
}

func (g *game) init() {
	g.Screen = gfx.NewScreen()
	g.World = world.NewWorld()
	g.Graphics = gfx.NewGraphics()
	g.InitInput()
	g.Gui = gui.New(0, 0, gfx.ScreenWidth, gfx.ScreenHeight)
	g.InitMenus()
	g.isPaused = true
}

func (g *game) Update(screen *ebiten.Image) error {

	g.Screen.SetScreen(screen)
	g.Gui.Update()
	g.Input.Update()

	if !g.isPaused {
		g.World.Update(g.tick)
		g.Screen.SetCameraCoords(g.World.PlayerUnit.GetPos())
		g.tick++
	}

	if !ebiten.IsDrawingSkipped() {
		if !g.isPaused {
			g.World.Draw(g.Screen)
		}
		g.Gui.Draw(screen)
		g.framesDrawn++
	}

	return nil
}

func (g *game) openContextMenu() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.rightMouseDown = true
	} else {
		if g.rightMouseDown {
			g.Gui.AddMenu("context", g.BuildContextMenu(ebiten.CursorPosition()))
		}
		g.rightMouseDown = false
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.rightMouseDown = false
		contextMenu := g.Gui.GetMenu("context")
		if contextMenu != nil {
			contextMenu.SetDisabled(true)
		}
	}
}

func (g *game) DebugInfo() string {
	aString := "Tick: " + strconv.Itoa(g.tick)
	aString += "\nFrame: " + strconv.Itoa(g.framesDrawn)
	if g.World.PlayerUnit != nil {
		x, y := g.World.PlayerUnit.GetPos()
		aString += "\nPlayer: " + strconv.Itoa(int(x)) + " / " + strconv.Itoa(int(y))
	}
	if g.debugMessage != "" {
		aString += "\n" + g.debugMessage
	}

	return aString
}

func (g *game) TogglePause() {
	if g.isPaused {
		g.UnPause()
		return
	}
	g.Pause()
}

func (g *game) UnPause() {
	if !g.World.IsInitialized() {
		g.World.Init()
	}
	g.isPaused = false
	g.Gui.SetDisabled("start", true)
	g.Gui.SetDisabled("context", true)
	g.Gui.SetDisabled("game", false)
}

func (g *game) Pause() {
	g.isPaused = true
	g.Gui.SetDisabled("context", true)
	g.Gui.SetDisabled("game", true)
	g.Gui.SetDisabled("start", false)
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
	g.Input.AddListener("RightPress", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERRIGHT, true)
	})
	g.Input.AddListener("RightRelease", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERRIGHT, false)
	})
	g.Input.AddListener("UpPress", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERUP, true)
	})
	g.Input.AddListener("UpRelease", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERUP, false)
	})
	g.Input.AddListener("DownPress", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERDOWN, true)
	})
	g.Input.AddListener("DownRelease", "updatePlayer", func(g *game) {
		g.UpdatePlayerMovement(world.PLAYERDOWN, false)
	})
}

func (g *game) UpdatePlayerMovement(dir world.PlayerDirection, value bool) {
	if g.World == nil || g.World.PlayerUnit == nil {
		return
	}
	g.World.PlayerUnit.SetMovement(dir, value)
}
