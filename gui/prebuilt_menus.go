package gui

import (
	"image"
	"log"
)

func StartMenu(gui *Gui) *menu {
	startMenu := newMenu(gui)
	startMenu.topPadding = 50

	buttonSize := image.Rect(0, 0, 150, 30)

	firstButton := NewButton(buttonSize, "Test")
	enablerButton := NewButton(buttonSize, "Enabler")
	hiddenButton := NewButton(buttonSize, "Hidden!")
	stopButton := NewButton(buttonSize, "Stop!")

	hiddenButton.disabled = true
	hiddenButton.onPressed = func(b *button) {
		aButton := NewButton(buttonSize, "Another Button")
		startMenu.addButton(aButton)
	}
	enablerButton.onPressed = func(b *button) {
		hiddenButton.disabled = !hiddenButton.disabled
	}
	stopButton.onPressed = func(b *button) {
		log.Fatal("Stopped")
	}

	startMenu.addButton(firstButton)
	startMenu.addButton(hiddenButton)
	startMenu.addButton(enablerButton)
	startMenu.addButton(stopButton)

	return startMenu
}
