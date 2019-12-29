package file

type DefClass interface {
	GetName() string
}

type SectorDef struct {
	Name          string
	CenterGraphic string
	Weight        int
	Land          int
}

func (d *SectorDef) GetName() string {
	return d.Name
}

type BuildingDef struct {
	Name    string
	Graphic string
}

func (d *BuildingDef) GetName() string {
	return d.Name
}
