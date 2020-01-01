package world

import (
	"github.com/pedritoelcabra/projectx/core/logger"
	"github.com/pedritoelcabra/projectx/world/tiling"
)

type node struct {
	location tiling.Coord
	f        float64
	g        float64
	h        float64
}

const DefaultPathMaxLength = 100

type PathOptions struct {
	MaxLength int
}

type Path struct {
	Path        []tiling.Coord
	Start       tiling.Coord
	End         tiling.Coord
	Nodes       int
	MaxLength   int
	Cost        float64
	Valid       bool
	CurrentStep int
}

func (p *Path) GetSteps() []tiling.Coord {
	return p.Path
}

func (p *Path) IsValid() bool {
	return p.Valid
}

func FindPath(start, end tiling.Coord) Path {
	options := PathOptions{DefaultPathMaxLength}
	aPath := FindPathWithOptions(start, end, options)
	if aPath.IsValid() {
		logMsg := "Found path: "
		for _, step := range aPath.GetSteps() {
			logMsg += step.ToString() + ", "
		}
		logger.General(logMsg, nil)
	} else {
		logger.General("Found no valid path from "+start.ToString()+" to "+end.ToString(), nil)
	}
	return aPath
}

func FindPathWithOptions(start, end tiling.Coord, options PathOptions) Path {
	aPath := Path{}
	aPath.MaxLength = options.MaxLength
	aPath.Valid = false
	aPath.Nodes = len(aPath.Path)
	aPath.Cost = 0.0
	aPath.Start = start
	aPath.End = end
	aPath.CurrentStep = 0
	return findValidPath(aPath)
}

func findValidPath(path Path) Path {

	return path
}
