package test

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/gui"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func SetWorkingDirectory() {
	dir, _ := os.Getwd()
	if !strings.Contains(dir, "test") {
		return
	}
	_ = os.Chdir(filepath.Join(dir, "/../.."))
}

func GetTextBox(menu *gui.Menu) *gui.TextBox {
	aBox := gui.NewTextBox()
	aBox.SetBox(image.Rect(0, 0, 400, 400))
	aBox.SetColor(color.White)
	aBox.SetText("test")
	aBox.SetTextSize(gui.FontSize12)
	aBox.SetHCentered(false)
	aBox.SetLeftPadding(10)
	menu.AddTextBox(aBox)
	return aBox
}

func GetGui() *gui.Gui {
	return gui.New(0, 0, 800, 600)
}

func GetMenu() *gui.Menu {
	aGui := GetGui()
	return gui.NewMenu(aGui)
}

func GetBox() image.Rectangle {
	return image.Rect(0, 0, 800, 600)
}

func GetText() string {
	return "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum"
}

func TestGui(t *testing.T) {
	SetWorkingDirectory()
	aGui := gui.New(0, 0, 800, 600)
	aGui.AddMenu("test", GetMenu())

	aImage, _ := ebiten.NewImage(800, 600, ebiten.FilterDefault)
	aGui.Draw(aImage)

	for i := 0; i < 1; i++ {
		aGui.Draw(aImage)
	}
}

func TestBuildTextBox(t *testing.T) {
	SetWorkingDirectory()
	aBox := GetTextBox(GetMenu())
	aBox.SetText(GetText())
	for i := 0; i < 1; i++ {
		aBox.BuildTextBoxImage(GetGui(), GetBox())
	}
}

func TestSplitText(t *testing.T) {
	SetWorkingDirectory()
	aGui := GetGui()
	aBox := GetTextBox(GetMenu())
	aBox.SetText(GetText())
	for i := 0; i < 1; i++ {
		aBox.InsertLineBreaks(aGui)
	}
}

func TestStringWidth(t *testing.T) {
	tt, _ := truetype.Parse(goregular.TTF)
	aFace := truetype.NewFace(tt, &truetype.Options{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	aBox := GetTextBox(GetMenu())
	for i := 0; i < 1000; i++ {
		aBox.EstimateStringBounds(aFace, GetText())
	}
}
