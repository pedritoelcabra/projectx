package gui

type Gui struct {
	menus []menu
}

func New() Gui {
	return Gui{}
}
