package world

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/world/tiling"
	"log"
)

type JobType int

const (
	GatherJob JobType = iota
	BuildJob
)

type Job struct {
	Worker   UnitPointer
	Building BuildingPointer
	Type     JobType
	Location tiling.Coord
}

func NewGatheringJob(building *Building, location tiling.Coord) *Job {
	aJob := &Job{}
	aJob.Worker = MakeEmptyUnitPointer()
	aJob.Type = GatherJob
	aJob.Location = location
	aJob.Building = building.GetPointer()
	return aJob
}

func NewBuildJob(building *Building) *Job {
	aJob := &Job{}
	aJob.Worker = MakeEmptyUnitPointer()
	aJob.Type = BuildJob
	aJob.Location = building.GetTile().GetCoord()
	aJob.Building = building.GetPointer()
	return aJob
}

func (j *Job) GetBuilding() *Building {
	return j.Building.Get()
}

func (j *Job) GetWorker() *Unit {
	return j.Worker.Get()
}

func (j *Job) GetLocation() tiling.Coord {
	return j.Location
}

func (j *Job) HireWorker(worker *Unit) {
	j.Worker = worker.GetPointer()
	worker.SetWork(j)
	if j.Building.Get() != nil {
		sector := j.Building.Get().GetSector()
		if sector != nil {
			sector.RemoveJob(j)
		}
	}
}

func (j *Job) Destroy() {
	if j.Worker.Get() != nil {
		j.Worker.Get().SetWork(nil)
	}
	if j.Building.Get() != nil {
		j.Building.Get().ClearJob()
		sector := j.Building.Get().GetSector()
		if sector != nil {
			sector.RemoveJob(j)
		}
	}
}

func (j *Job) GetName() string {
	switch j.Type {
	case BuildJob:
		return "Building Job"
	case GatherJob:
		return "Gathering Job"
	}
	return "Unknown Job"
}

func (j *Job) AddWork() {
	building := j.GetBuilding()
	switch j.Type {
	case BuildJob:
		building.AddConstructionProgress(1)
	case GatherJob:
		if building.Get(GatherStatus) != GatherStatusAvailable {
			j.Destroy()
			return
		}
		tile := theWorld.Grid.Tile(j.GetLocation())
		if tile == nil {
			log.Fatal("no target for gathering")
		}
		amount := tile.GetResourceAmount()
		newAmount := amount - 1
		tile.SetResourceAmount(newAmount)
		sector := tile.GetSector()
		if sector == nil {
			log.Fatal("target has no sector!")
		}
		sector.GetInventory().AddItem(defs.GetMaterialDef(building.Template.Gathers).ID, 1)
		if newAmount <= 0 {
			j.Destroy()
		}
	}
}
