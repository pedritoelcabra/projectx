package core

import (
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/gui"
	"image"
	"image/color"
)

const (
	EntityMenuWidth         = 200
	EntityMenuBottomPadding = 100
	ClickableDistance       = 20
	TitleSize               = 36
)

type Entity interface {
	GetName() string
}

func (g *game) BuildEntityMenu() *gui.Menu {
	menu := gui.NewMenu(g.Gui)
	menu.SetDisabled(true)
	menu.SetBG(color.Black)
	return menu
}

func (g *game) OpenEntityMenu(x, y int) {
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

func (g *game) ShowEntity(entity Entity) {
	menu := g.BuildEntityMenu()
	menu.SetDisabled(false)
	AddEntityTitle(menu, entity.GetName())
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
