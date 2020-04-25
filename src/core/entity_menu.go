package core

import (
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/gui"
	"github.com/pedritoelcabra/projectx/src/world"
	"image"
	"image/color"
)

const (
	EntityMenuWidth         = 200
	EntityMenuBottomPadding = 100
	ClickableDistance       = 20
	TitleSize               = 36
)

func (g *game) BuildEntityMenu() *gui.Menu {
	menu := gui.NewMenu(g.Gui)
	menu.SetDisabled(true)
	menu.SetBG(color.Black)
	return menu
}

func (g *game) OpenEntityMenu(x, y int) {
	clickedUnit := g.World.ClosestUnitWithinRadius(x, y, ClickableDistance)
	if clickedUnit != nil {
		g.ShowUnitEntity(clickedUnit)
		return
	}
	mouseCoord := g.MouseTileCoord()
	tile := g.World.Grid.Tile(mouseCoord)
	building := tile.Building.Get()
	if building != nil {
		g.ShowBuildingEntity(building)
	}
}

func (g *game) ShowUnitEntity(unit *world.Unit) {
	menu := g.BuildEntityMenu()
	menu.SetDisabled(false)
	AddEntityTitle(menu, unit.GetName())
	menu.SetLeftPadding(gfx.ScreenWidth - EntityMenuWidth)
	menu.SetBottomPadding(EntityMenuBottomPadding)
	g.Gui.AddMenu(EntityMenu, menu)
}

func (g *game) ShowBuildingEntity(building *world.Building) {
	menu := g.BuildEntityMenu()
	menu.SetDisabled(false)
	AddEntityTitle(menu, building.GetName())
	menu.SetLeftPadding(gfx.ScreenWidth - EntityMenuWidth)
	menu.SetBottomPadding(EntityMenuBottomPadding)
	g.Gui.AddMenu(EntityMenu, menu)
}

func AddEntityTitle(menu *gui.Menu, title string) {
	aBox := gui.NewTextBox()
	aBox.SetBox(image.Rect(0, 0, EntityMenuWidth, gfx.ScreenHeight-EntityMenuBottomPadding))
	aBox.SetColor(color.White)
	aBox.SetText(title)
	aBox.SetTextSize(gui.FontSize20)
	aBox.SetHCentered(true)
	menu.AddTextBox(aBox)
}
