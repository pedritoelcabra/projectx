package gui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"strings"
)

type TextBox struct {
	box          image.Rectangle
	drawBox      image.Rectangle
	text         string
	textSize     FontSize
	hasDrawnText bool
	contentBuf   *ebiten.Image
	vCenter      bool
	hCenter      bool
	leftPadding  int
	topPadding   int
	fontColor    color.Gray16
	OnUpdate     func(t *TextBox)
}

func NewTextBox() *TextBox {
	aTextBox := &TextBox{}
	aTextBox.textSize = FontSize12
	aTextBox.hasDrawnText = false
	return aTextBox
}

func (t *TextBox) SetBox(box image.Rectangle) {
	t.box = box
}

func (t *TextBox) SetLeftPadding(value int) {
	t.leftPadding = value
}

func (t *TextBox) SetTopPadding(value int) {
	t.topPadding = value
}

func (t *TextBox) SetHCentered(centered bool) {
	t.hCenter = centered
}

func (t *TextBox) SetVCentered(centered bool) {
	t.vCenter = centered
}

func (t *TextBox) SetColor(color color.Gray16) {
	t.fontColor = color
}

func (t *TextBox) SetText(text string) {
	if text == t.text {
		return
	}
	t.text = text
	t.hasDrawnText = false
}

func (t *TextBox) SetTextSize(size int) {
	t.textSize = FontSize(size)
}

func (t *TextBox) BuildTextBoxImage(gui *Gui, box image.Rectangle) {

	offsetDrawBox(&t.drawBox, &box, &t.box)
	w := t.getWidth()
	h := t.getHeight()
	if t.contentBuf == nil {
		t.contentBuf, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	}

	t.contentBuf.Clear()

	maxWidth := 0
	maxHeight := 0
	for i, line := range strings.Split(t.text, "\n") {
		x := 0
		y := 0 + i*lineHeight + lineHeight - (lineHeight-gui.uiFontHeights[t.textSize])/2
		if y < -lineHeight {
			continue
		}

		currentBounds, _ := font.BoundString(gui.uiFonts[t.textSize], line)
		currentWidth := currentBounds.Max.X.Ceil()
		currentHeight := -currentBounds.Min.Y.Ceil()

		if currentHeight > maxHeight {
			maxHeight = currentHeight
		}
		if currentWidth > maxWidth {
			maxWidth = currentWidth
		}

		text.Draw(t.contentBuf, line, gui.uiFonts[t.textSize], x, y, t.fontColor)
	}

	if t.vCenter && maxHeight < h {
		t.drawBox.Min.Y += (h - maxHeight) / 2
	}

	if t.hCenter && maxWidth < w {
		t.drawBox.Min.X += (w - maxWidth) / 2
	}

	t.drawBox.Min.X += t.leftPadding
	t.drawBox.Min.Y += t.topPadding
	t.hasDrawnText = true
}
