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

func MakeEmptyBuildingPointer() BuildingPointer {
	return MakeBuildingPointer(-1)
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

type UnitPointer struct {
	Id          UnitKey
	building    *Unit
	initialised bool
}

func MakeUnitPointer(key UnitKey) UnitPointer {
	aPointer := UnitPointer{}
	aPointer.Id = key
	return aPointer
}

func MakeEmptyUnitPointer() UnitPointer {
	return MakeUnitPointer(-1)
}

func (up *UnitPointer) Get() *Unit {
	if up.Id < 0 {
		return nil
	}
	if !up.initialised {
		up.building = theWorld.GetUnit(up.Id)
		up.initialised = true
	}
	return up.building
}
