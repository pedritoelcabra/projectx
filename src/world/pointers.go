package world

type BuildingPointer struct {
	Id          BuildingKey
	building    *Building
	initialised bool
}

func MakeBuildingPointer(key BuildingKey) BuildingPointer {
	aPointer := BuildingPointer{}
	aPointer.Id = key
	return aPointer
}

func (bp *BuildingPointer) Get() *Building {
	if bp.Id < 0 {
		return nil
	}
	if !bp.initialised {
		bp.building = theWorld.GetBuilding(bp.Id)
		bp.initialised = true
	}
	return bp.building
}
