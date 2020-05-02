package core

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/logger"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/gui"
	"image"
	"image/color"
)

func (g *game) BuildBuildings() *gui.Menu {
	buildingMenu := gui.NewMenu(g.Gui)
	buildingMenu.SetHorizontalMenu(true)
	buildingMenu.SetBG(color.Black)

	titleSize := 100
	aBox := gui.NewTextBox()
	aBox.SetBox(image.Rect(0, 0, gfx.ScreenWidth, titleSize))
	aBox.SetColor(color.White)
	aBox.SetText("Buildings")
	aBox.SetTextSize(gui.FontSize24)
	aBox.SetHCentered(true)
	buildingMenu.AddTextBox(aBox)

	buttonSize := image.Rect(0, 0, 100, 100)
	for _, def := range defs.BuildingDefs() {
		buildingName := def.Name
		buildingDef := def
		buildingButton := gui.NewButton(buttonSize, buildingName)
		buildingButton.OnPressed = func(b *gui.Button) {
			g.SelectBuilding(buildingDef)
		}
		buildingButton.SetImage(gfx.GetImage(gfx.GetSpriteKey(def.Graphic)))
		buildingButton.SetVCentered(false)
		buildingMenu.AddButton(buildingButton)
	}

	buildingMenu.SetDisabled(true)
	return buildingMenu
}

func (g *game) SelectBuilding(building *defs.BuildingDef) {
	buildingName := building.Name
	logger.General(buildingName+" selected", nil)
	g.Gui.SetDisabled(BuildingMenu, true)
	g.ShowEntity(building)
	g.PlacementManager.SetBuilding(building)
}
