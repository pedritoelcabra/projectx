package defs

type DefClass interface {
	GetName() string
	GetGraphic() string
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

func (d *SectorDef) GetGraphic() string {
	return d.CenterGraphic
}

type BuildingDef struct {
	Name    string
	Graphic string
}

func (d *BuildingDef) GetName() string {
	return d.Name
}

func (d *BuildingDef) GetGraphic() string {
	return d.Graphic
}

var Definitions = make(map[string]map[string]DefClass)

func GetDefs(defType string) map[string]DefClass {
	return Definitions[defType]
}
