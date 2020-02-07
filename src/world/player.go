package world

import (
	"github.com/pedritoelcabra/projectx/src/core/logger"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"strconv"
)

type PlayerDirection int

const (
	PLAYERUP PlayerDirection = iota
	PLAYERLEFT
	PLAYERDOWN
	PLAYERRIGHT
)

type Player struct {
	unit            *Unit
	MovingUp        bool
	MovingLeft      bool
	MovingDown      bool
	MovingRight     bool
	RespawnCooldown int
}

func NewPlayer() *Player {
	aPlayer := &Player{}
	aPlayer.Init()
	aPlayer.unit = NewUnit("Player", tiling.NewCoord(0, 0))
	aPlayer.unit.Name = "You"
	return aPlayer
}

func (p *Player) Init() {
	p.unit = theWorld.GetUnit(0)
}

func (p *Player) DrawSprite(screen *gfx.Screen) {
	p.unit.DrawSprite(screen)
}

func (p *Player) SetPosition(x, y float64) {
	p.unit.SetPosition(x, y)
}

func (p *Player) MoveToHomeSector() {
	sector := theWorld.GetSector(SectorKey(p.unit.GetF(HomeSector)))
	if sector != nil {
		sectorCenterTile := theWorld.Grid.Tile(sector.GetCenter())
		p.unit.SetPosition(sectorCenterTile.GetF(RenderX), sectorCenterTile.GetF(RenderY)+100)
	}
}

func (p *Player) Update() {
	p.CheckForPlayerDeath()
	if p.unit.IsAlive() {
		p.UpdateDestination()
	}
}

func (p *Player) CheckForPlayerDeath() {
	if p.unit.IsAlive() {
		return
	}
	if p.RespawnCooldown < 0 {
		p.RespawnCooldown = 180
	}
	if p.RespawnCooldown >= 0 {
		p.RespawnCooldown--
		if (p.RespawnCooldown % 10) == 0 {
			logger.General("Respawning in "+strconv.Itoa(p.RespawnCooldown), nil)
		}
	}
	if p.RespawnCooldown == 0 {
		p.unit.Alive = true
		p.unit.SetToMaxHealth()
		p.MoveToHomeSector()
		p.RespawnCooldown = -1
	}
}

func (p *Player) UpdateDestination() {
	destX := p.unit.X
	destY := p.unit.Y
	displacement := float64(1000)
	if p.MovingUp && !p.MovingDown {
		destY -= displacement
	}
	if p.MovingDown && !p.MovingUp {
		destY += displacement
	}
	if p.MovingLeft && !p.MovingRight {
		destX -= displacement
	}
	if p.MovingRight && !p.MovingLeft {
		destX += displacement
	}
	p.unit.SetDestination(destX, destY)
}

func (p *Player) GetPos() (x, y float64) {
	return p.unit.GetPos()
}

func (p *Player) GetX() float64 {
	return p.unit.GetX()
}

func (p *Player) GetY() float64 {
	return p.unit.GetY()
}

func (p *Player) SetMovement(direction PlayerDirection, value bool) {
	switch direction {
	case PLAYERUP:
		p.MovingUp = value
	case PLAYERLEFT:
		p.MovingLeft = value
	case PLAYERDOWN:
		p.MovingDown = value
	case PLAYERRIGHT:
		p.MovingRight = value
	}
}
