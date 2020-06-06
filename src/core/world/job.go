package world

import "github.com/pedritoelcabra/projectx/src/core/world/tiling"

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
