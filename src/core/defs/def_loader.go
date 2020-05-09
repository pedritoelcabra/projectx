package defs

const defFolder = "defs/"

func InitDefs() {
	LoadBuildingDefs()
	LoadSectorDefs()
	LoadResourceDefs()
	LoadUnitDefs()
	LoadFactionDefs()
	LoadMaterialDefs()
}
