package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/core/file"
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/gui"
	"github.com/pedritoelcabra/projectx/world"
	"github.com/pedritoelcabra/projectx/world/grid"
	"math/rand"
	"strconv"
	"time"
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
	if g.HasLoadedWorld() {
		x, y := g.World.PlayerUnit.GetPos()
		aString += "\nPlayer Pos: " + strconv.Itoa(int(x)) + " / " + strconv.Itoa(int(y))
		tx, ty := world.PosToTile(int(x), int(y))
		aString += "\nPlayer Tile: " + strconv.Itoa(tx) + " / " + strconv.Itoa(ty)
		tile := g.World.Grid.Tile(grid.NewCoord(tx, ty))
		height := tile.Get(grid.Height)
		aString += "\nTile Height: " + strconv.Itoa(height)

		mx, my := ebiten.CursorPosition()
		cx, cy := g.Screen.GetCameraCoords()
		mx += int(cx)
		my += int(cy)
		aString += "\nMouse Pos: " + strconv.Itoa(int(mx)) + " / " + strconv.Itoa(int(my))
		mtx, mty := world.PosToTile(int(mx), int(my))
		aString += "\nMouse Tile: " + strconv.Itoa(mtx) + " / " + strconv.Itoa(mty)
		mTile := g.World.Grid.Tile(grid.NewCoord(mtx, mty))
		mHeight := mTile.Get(grid.Height)
		aString += "\nMouse Tile Height: " + strconv.Itoa(mHeight)

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
	g.isPaused = false
	g.Gui.SetDisabled("start", true)
	g.Gui.SetDisabled("context", true)
	g.Gui.SetDisabled("game", false)
}

func (g *game) HasLoadedWorld() bool {
	return g.World != nil && g.World.IsInitialized()
}

func (g *game) InitializeNewWorld() {
	g.World = world.NewWorld()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	g.World.SetSeed(r1.Intn(10000))
	g.World.Init()
	g.InitMenus()
	g.UnPause()
}

func (g *game) LoadSavedWorld(data file.SaveGameData) {
	g.World = world.NewWorld()
	g.World.SetSeed(data.Seed)
	g.World.Init()
	g.InitMenus()
	g.UnPause()
}

func (g *game) Pause() {
	g.isPaused = true
	g.Gui.SetDisabled("context", true)
	g.Gui.SetDisabled("game", true)
	g.Gui.SetDisabled("start", false)
}

func (g *game) UpdatePlayerMovement(dir world.PlayerDirection, value bool) {
	if g.World == nil || g.World.PlayerUnit == nil {
		return
	}
	g.World.PlayerUnit.SetMovement(dir, value)
}

func (g *game) SaveGameState() {
	state := file.SaveGameData{}
	state.Seed = g.World.GetSeed()
	state.Tick = g.tick
	state.Player = *g.World.PlayerUnit
	file.SaveToFile(state, file.DefaultSaveGameName)
}

func (g *game) LoadGameState() {
	dataStructure := file.LoadFromFile(file.DefaultSaveGameName)
	g.tick = dataStructure.Tick
	g.LoadSavedWorld(dataStructure)
}
