package gui

var leftClickAvailable bool
var rightClickAvailable bool

func LeftClickAvailable() bool {
	return leftClickAvailable
}

func RightClickAvailable() bool {
	return rightClickAvailable
}

func UseRightClick() {
	rightClickAvailable = false
}

func UseLeftClick() {
	leftClickAvailable = false
}

func ResetClicks() {
	leftClickAvailable = true
	rightClickAvailable = true
}
