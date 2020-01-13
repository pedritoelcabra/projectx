package defs

const defFolder = "defs/"

func InitDefs() {
	LoadBuildingDefs()
	LoadSectorDefs()
	LoadVegetationDefs()
	LoadUnitDefs()
}
