package gui

import (
	"github.com/hajimehoshi/ebiten"
	"image"
)

type drawable interface {
	draw(*Gui, image.Rectangle)
	update()
	getWidth() int
	getHeight() int
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////

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

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////

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

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////

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

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////

func (m *Menu) draw(gui *Gui, box image.Rectangle) {
	if m.disabled {
		return
	}
	drawSpace := box
	drawSpace.Min.Y += m.topPadding
	drawSpace.Min.X += m.leftPadding
	drawSpace.Min.X += m.centeredOffset(box)
	drawSpace.Max.X -= m.rightPadding
	drawSpace.Max.Y -= m.bottomPadding
	currentBatchMaxDimension := 0
	if m.hasBG {
		bg, _ := ebiten.NewImage(drawSpace.Dx(), drawSpace.Dy(), ebiten.FilterNearest)
		_ = bg.Fill(m.background)
		opts := &ebiten.DrawImageOptions{}
		gui.drawImage(bg, drawSpace, opts)
	}
	for _, component := range m.components {
		width := component.getWidth()
		height := component.getHeight()
		if m.horizontalMenu {
			if drawSpace.Min.X+width >= box.Max.X {
				drawSpace.Min.X = box.Min.X + m.leftPadding + m.centeredOffset(box)
				drawSpace.Min.Y += currentBatchMaxDimension
			}
			if currentBatchMaxDimension < height {
				currentBatchMaxDimension = height
			}
		} else {
			if drawSpace.Min.Y+height >= box.Max.Y {
				drawSpace.Min.Y = box.Min.Y + m.topPadding
				drawSpace.Min.Y += currentBatchMaxDimension
			}
			if currentBatchMaxDimension < width {
				currentBatchMaxDimension = width
			}
		}
		component.draw(gui, drawSpace)
		if m.horizontalMenu {
			drawSpace.Min.X += width
		} else {
			drawSpace.Min.Y += height
		}
	}
}

func (b *Button) draw(gui *Gui, box image.Rectangle) {
	if b.disabled {
		return
	}
	offsetDrawBox(&b.drawBox, &box, &b.box)
	gui.draw(b.drawBox, imageSrcRects[b.buttonImage])
	if b.itemImage != nil {
		op := &ebiten.DrawImageOptions{}
		gui.drawImage(b.itemImage, b.drawBox, op)
	}
	b.textBoxImg.draw(gui, box)
}

func (t *TextBox) draw(gui *Gui, box image.Rectangle) {
	if !t.hasDrawnText {
		t.BuildTextBoxImage(gui, box)
	}
	op := &ebiten.DrawImageOptions{}
	gui.drawImage(t.contentBuf, t.drawBox, op)
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////

func offsetDrawBox(d *image.Rectangle, p *image.Rectangle, b *image.Rectangle) {
	d.Min.X = p.Min.X + b.Min.X
	d.Max.X = p.Min.X + b.Max.X
	d.Min.Y = p.Min.Y + b.Min.Y
	d.Max.Y = p.Min.Y + b.Max.Y

}
