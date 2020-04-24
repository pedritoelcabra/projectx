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
	TitleSize               = 36
)

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
