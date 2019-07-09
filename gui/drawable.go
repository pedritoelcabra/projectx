package gui

import "image"

type drawable interface {
	draw(drawFunction, image.Rectangle)
	getWidth() int
	getHeight() int
}

func (m *menu) getWidth() int {
	maxWidth := 0
	for _, component := range m.components {
		if component.getWidth() > maxWidth {
			maxWidth = component.getWidth()
		}
	}
	return maxWidth
}

func (b *button) getWidth() int {
	return b.box.Max.X - b.box.Min.X
}

func (m *menu) getHeight() int {
	height := 0
	for _, component := range m.components {
		height += component.getHeight()
	}
	return height
}

func (b *button) getHeight() int {
	return b.box.Max.Y - b.box.Min.Y
}

func (m *menu) draw(drawFun drawFunction, box image.Rectangle) {
	drawSpace := box
	drawSpace.Min.Y += m.topPadding
	drawSpace.Min.X += m.leftPadding
	drawSpace.Min.X += m.centeredOffset(box)
	for _, component := range m.components {
		component.draw(drawFun, drawSpace)
		drawSpace.Min.Y += component.getHeight()
	}
}

func (b *button) draw(drawFun drawFunction, box image.Rectangle) {
	b.drawBox.Min.X = box.Min.X + b.box.Min.X
	b.drawBox.Max.X = box.Min.X + b.box.Max.X
	b.drawBox.Min.Y = box.Min.Y + b.box.Min.Y
	b.drawBox.Max.Y = box.Min.Y + b.box.Max.Y
	drawFun(b.drawBox, imageSrcRects[imageTypeButton])
}
