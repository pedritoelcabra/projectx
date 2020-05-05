package gui

import "image"

type Entity interface {
	GetName() string
	GetStats() string
	AddButtonsToEntityMenu(menu *Menu, size image.Rectangle)
}
