package world

import (
	"github.com/pedritoelcabra/projectx/core/logger"
	"github.com/pedritoelcabra/projectx/world/tiling"
)

type node struct {
	location  tiling.Coord
	parent    *node
	total     float64
	current   float64
	predicted float64
}

const DefaultPathMaxLength = 30
const DefaultPathMaxNodes = 300

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
	open        []*node
	closed      []*node
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
	aPath.findValidPath()
	return aPath
}

func (p *Path) findValidPath() {
	startNode := p.NewNode(p.Start, nil)
	p.open = append(p.open, startNode)
	for !p.Valid && len(p.closed) < DefaultPathMaxNodes {
		bestNode := p.bestOpenNode()
		if bestNode == nil {
			break
		}
		p.ProcessNewNode(bestNode)
	}
	p.generateCleanPath(p.closed)
	p.open = []*node{}
	p.closed = []*node{}
}

func (p *Path) generateCleanPath(closed []*node) {
	if !p.Valid {
		return
	}
	for _, node := range closed {
		p.Path = append(p.Path, node.location)
	}
}

func (p *Path) ProcessNewNode(parent *node) {
	p.closed = append(p.closed, parent)
	neighbours := theWorld.Grid.Tile(parent.location).Neighbours()
	for _, tile := range neighbours {
		if tile.IsImpassable() {
			continue
		}
		if p.ExistsInList(tile.GetCoord()) {
			continue
		}
		newNode := p.NewNode(tile.GetCoord(), parent)
		p.open = append(p.open, newNode)
		if tile.Coord().Equals(p.End) {
			p.closed = append(p.closed, newNode)
			p.Valid = true
			p.MaxLength = len(p.closed)
			p.Cost = newNode.current
			return
		}
	}
}

func (p *Path) ExistsInList(aCoord tiling.Coord) bool {
	for _, bNode := range p.open {
		if aCoord.Equals(bNode.location) {
			return true
		}
	}
	return false
}

func (p *Path) bestOpenNode() *node {
	bestKey := 0
	bestScore := 99999999.0
	for key, node := range p.open {
		if bestScore > node.total {
			bestScore = node.total
			bestKey = key
		}
	}
	aNode := p.open[bestKey]
	p.open[bestKey] = p.open[len(p.open)-1]
	p.open = p.open[:len(p.open)-1]
	return aNode
}

func (p *Path) NewNode(location tiling.Coord, parent *node) *node {
	aNode := &node{}
	aNode.location = location
	aNode.parent = parent
	aNode.current = getTileCost(location)
	if parent != nil {
		aNode.current += parent.current
	}
	aNode.predicted = tiling.HexDistance(location, p.End)
	aNode.total = aNode.current + aNode.predicted
	return aNode
}

func getTileCost(coord tiling.Coord) float64 {
	return theWorld.Grid.Tile(coord).GetF(MovementCost)
}
