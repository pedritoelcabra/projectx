package gui

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"log"
)

const (
	lineHeight = 16
)

type Gui struct {
	box   image.Rectangle
	menus map[string]*Menu

	uiImage       *ebiten.Image
	screen        *ebiten.Image
	uiFont        font.Face
	uiFontMHeight int
}

func New(x, y, w, h int) *Gui {
	aGui := &Gui{}
	aGui.box = image.Rect(x, y, w, h)
	aGui.menus = make(map[string]*Menu)

	img, _, err := ebitenutil.NewImageFromFile("gui/ui.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	aGui.uiImage = img

	tt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
	aGui.uiFont = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	b, _, _ := aGui.uiFont.GlyphBounds('M')
	aGui.uiFontMHeight = (b.Max.Y - b.Min.Y).Ceil()

	return aGui
}

func (g *Gui) AddMenu(key string, menu *Menu) {
	g.menus[key] = menu
}

func (g *Gui) GetMenu(key string) *Menu {
	return g.menus[key]
}

func (g *Gui) SetDisabled(key string, value bool) {
	aMenu := g.menus[key]
	if aMenu != nil {
		aMenu.SetDisabled(value)
	}
}

func (g *Gui) ToggleDisabled(key string) {
	aMenu := g.menus[key]
	if aMenu != nil {
		aMenu.SetDisabled(!aMenu.disabled)
	}
}

func (g *Gui) DisableAllMenus() {
	for key, menu := range g.menus {
		if key == "debug" {
			continue
		}
		menu.SetDisabled(true)
	}
}

func (g *Gui) Update() {
	for _, menu := range g.menus {
		menu.update()
	}
}

func (g *Gui) Draw(screen *ebiten.Image) {
	g.screen = screen
	for _, menu := range g.menus {
		menu.draw(g, g.box)
	}
}

func (g *Gui) draw(dstRect image.Rectangle, srcRect image.Rectangle) {
	srcX := srcRect.Min.X
	srcY := srcRect.Min.Y
	srcW := srcRect.Dx()
	srcH := srcRect.Dy()

	dstX := dstRect.Min.X
	dstY := dstRect.Min.Y
	dstW := dstRect.Dx()
	dstH := dstRect.Dy()

	op := &ebiten.DrawImageOptions{}
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			op.GeoM.Reset()

			sx := srcX
			sy := srcY
			sw := srcW / 4
			sh := srcH / 4
			dx := 0
			dy := 0
			dw := sw
			dh := sh
			switch i {
			case 1:
				sx = srcX + srcW/4
				sw = srcW / 2
				dx = srcW / 4
				dw = dstW - 2*srcW/4
			case 2:
				sx = srcX + 3*srcW/4
				dx = dstW - srcW/4
			}
			switch j {
			case 1:
				sy = srcY + srcH/4
				sh = srcH / 2
				dy = srcH / 4
				dh = dstH - 2*srcH/4
			case 2:
				sy = srcY + 3*srcH/4
				dy = dstH - srcH/4
			}

			op.GeoM.Scale(float64(dw)/float64(sw), float64(dh)/float64(sh))
			op.GeoM.Translate(float64(dx), float64(dy))
			op.GeoM.Translate(float64(dstX), float64(dstY))
			g.screen.DrawImage(g.uiImage.SubImage(image.Rect(sx, sy, sx+sw, sy+sh)).(*ebiten.Image), op)
		}
	}
}

func (g *Gui) drawImage(image *ebiten.Image, destRect image.Rectangle, op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(float64(destRect.Min.X), float64(destRect.Min.Y))
	g.screen.DrawImage(image, op)
}
