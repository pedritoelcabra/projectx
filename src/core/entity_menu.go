package core

import (
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/gui"
	"image"
	"image/color"
)

const (
	EntityMenuWidth         = 200
	EntityMenuHeight        = gfx.ScreenHeight - 200
	EntityMenuTitleHeight   = 100
	EntityMenuBodyHeight    = EntityMenuHeight - EntityMenuTitleHeight
	EntityMenuBottomPadding = 100
	ClickableDistance       = 20
)

func (g *game) BuildEntityMenu() *gui.Menu {
	menu := gui.NewMenu(g.Gui)
	menu.SetDisabled(true)
	menu.SetBG(color.Black)
	return menu
}

func (g *game) OpenEntityMenu(x, y int) {
	if g.isPaused {
		return
	}
	clickedUnit := g.World.ClosestUnitWithinRadius(x, y, ClickableDistance)
	if clickedUnit != nil {
		g.ShowEntity(clickedUnit)
		return
	}
	mouseCoord := g.MouseTileCoord()
	tile := g.World.Grid.Tile(mouseCoord)
	building := tile.Building.Get()
	if building != nil {
		g.ShowEntity(building)
	}
}

func (g *game) DisableEntityMenu() {
	g.Gui.SetDisabled(EntityMenu, true)
	g.World.SetDisplayEntity(nil)
	g.PlacementManager.UnSetBuilding()
}

func (g *game) ShowEntity(entity gui.Entity) {
	g.World.SetDisplayEntity(entity)
	menu := g.BuildEntityMenu()
	menu.SetDisabled(false)
	AddEntityTitle(menu, entity)
	AddEntityText(menu, entity)
	menu.SetLeftPadding(gfx.ScreenWidth - EntityMenuWidth)
	menu.SetBottomPadding(EntityMenuBottomPadding)
	buttonSize := image.Rect(0, 0, 150, 30)
	entity.AddButtonsToEntityMenu(menu, buttonSize)
	g.Gui.AddMenu(EntityMenu, menu)
}

func GetEntityText(entity gui.Entity) string {
	text := "\n" + entity.GetStats()
	return text
}

func AddEntityText(menu *gui.Menu, entity gui.Entity) {
	aBox := gui.NewTextBox()
	aBox.SetBox(image.Rect(0, 0, EntityMenuWidth, EntityMenuBodyHeight))
	aBox.SetColor(color.White)
	aBox.SetText(GetEntityText(entity))
	aBox.SetTextSize(gui.FontSize12)
	aBox.SetHCentered(false)
	aBox.SetLeftPadding(10)
	aBox.OnUpdate = func(t *gui.TextBox) {
		t.SetText(GetEntityText(entity))
	}
	menu.AddTextBox(aBox)
}

func AddEntityTitle(menu *gui.Menu, entity gui.Entity) {
	aBox := gui.NewTextBox()
	aBox.SetBox(image.Rect(0, 0, EntityMenuWidth, EntityMenuTitleHeight))
	aBox.SetColor(color.White)
	aBox.SetText(entity.GetName())
	aBox.SetTextSize(gui.FontSize20)
	aBox.SetHCentered(true)
	aBox.SetTopPadding(30)
	menu.AddTextBox(aBox)
}
