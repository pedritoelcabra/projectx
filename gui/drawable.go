package gui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image"
	"strings"
)

type drawable interface {
	draw(*Gui, image.Rectangle)
	update()
	getWidth() int
	getHeight() int
}

func (m *Menu) update() {
	if m.disabled {
		return
	}
	for _, component := range m.components {
		component.update()
	}
}

func (b *Button) update() {
	if b.disabled {
		return
	}
	mouseOverButton := b.mouseIsOverButton()
	if mouseOverButton {
		b.buttonImage = imageTypeButtonPressed
	} else {
		b.buttonImage = imageTypeButton
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if mouseOverButton {
			b.mouseDown = true
		} else {
			b.mouseDown = false
		}
	} else {
		if b.mouseDown && b.OnPressed != nil {
			b.OnPressed(b)
		}
		b.mouseDown = false
	}
}

func (t *TextBox) update() {
	if t.OnUpdate == nil {
		return
	}
	t.OnUpdate(t)
}

func (m *Menu) getWidth() int {
	maxWidth := 0
	for _, component := range m.components {
		if component.getWidth() > maxWidth {
			maxWidth = component.getWidth()
		}
	}
	return maxWidth
}

func (b *Button) getWidth() int {
	if b.disabled {
		return 0
	}
	return b.box.Max.X - b.box.Min.X
}

func (t *TextBox) getWidth() int {
	return t.box.Max.X - t.box.Min.X
}

func (m *Menu) getHeight() int {
	height := 0
	for _, component := range m.components {
		height += component.getHeight()
	}
	return height
}

func (b *Button) getHeight() int {
	if b.disabled {
		return 0
	}
	return b.box.Max.Y - b.box.Min.Y
}

func (t *TextBox) getHeight() int {
	return t.box.Max.Y - t.box.Min.Y
}

func (m *Menu) draw(gui *Gui, box image.Rectangle) {
	if m.disabled {
		return
	}
	drawSpace := box
	drawSpace.Min.Y += m.topPadding
	drawSpace.Min.X += m.leftPadding
	drawSpace.Min.X += m.centeredOffset(box)
	for _, component := range m.components {
		component.draw(gui, drawSpace)
		drawSpace.Min.Y += component.getHeight()
	}
}

func (b *Button) draw(gui *Gui, box image.Rectangle) {
	if b.disabled {
		return
	}
	offsetDrawBox(&b.drawBox, &box, &b.box)
	gui.draw(b.drawBox, imageSrcRects[b.buttonImage])
	b.textBoxImg.draw(gui, box)
}

func (t *TextBox) draw(gui *Gui, box image.Rectangle) {
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
		y := 0 + i*lineHeight + lineHeight - (lineHeight-gui.uiFontMHeight)/2
		if y < -lineHeight {
			continue
		}

		currentBounds, _ := font.BoundString(gui.uiFont, line)
		currentWidth := currentBounds.Max.X.Ceil()
		currentHeight := -currentBounds.Min.Y.Ceil()

		if currentHeight > maxHeight {
			maxHeight = currentHeight
		}
		if currentWidth > maxWidth {
			maxWidth = currentWidth
		}

		text.Draw(t.contentBuf, line, gui.uiFont, x, y, t.fontColor)
	}

	if t.vCenter && maxHeight < h {
		t.drawBox.Min.Y += (h - maxHeight) / 2
	}

	if t.hCenter && maxWidth < w {
		t.drawBox.Min.X += (w - maxWidth) / 2
	}

	t.drawBox.Min.X += t.leftPadding
	t.drawBox.Min.Y += t.topPadding

	op := &ebiten.DrawImageOptions{}
	gui.drawImage(t.contentBuf, t.drawBox, op)
}

func offsetDrawBox(d *image.Rectangle, p *image.Rectangle, b *image.Rectangle) {
	d.Min.X = p.Min.X + b.Min.X
	d.Max.X = p.Min.X + b.Max.X
	d.Min.Y = p.Min.Y + b.Min.Y
	d.Max.Y = p.Min.Y + b.Max.Y

}
