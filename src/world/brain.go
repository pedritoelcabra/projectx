package world

import (
	"github.com/pedritoelcabra/projectx/src/core/logger"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"log"
	"strconv"
)

type Brain struct {
	owner           *Unit
	Type            int
	CurrentState    int
	LastUpdated     int
	TargetKey       UnitKey
	BusyTime        int
	UpdateFrequency int
	target          *Unit
}

func NewBrain() *Brain {
	aBrain := &Brain{}
	aBrain.CurrentState = Idle
	aBrain.TargetKey = 0
	aBrain.LastUpdated = 0
	aBrain.BusyTime = 0
	aBrain.UpdateFrequency = 1
	return aBrain
}

const (
	Idle = iota
	Chase
	Flee
	Attack
	Patrol

	IdleMoveDistance = 10
	IdleMoveChance   = 25
)

func (b *Brain) ProcessState() {
	if !b.NeedsUpdating() {
		return
	}
	b.SetUpdateFrequency()
	logger.General("Updating unit "+strconv.Itoa(int(b.owner.Id))+" on tick "+strconv.Itoa(theWorld.GetTick()), nil)
	b.LastUpdated = theWorld.GetTick()
	b.ResolveState()
	switch b.CurrentState {
	case Idle:
		b.Idle()
		return
	case Chase:
		b.Chase()
		return
	case Attack:
		b.Attack()
		return
	}
	log.Fatal("Unknown state: " + strconv.Itoa(b.CurrentState))
}

func (b *Brain) SetUpdateFrequency() {
	distanceToPlayer := b.DistanceToUnit(theWorld.PlayerUnit.unit)
	if distanceToPlayer <= gfx.ScreenWidth {
		b.UpdateFrequency = 10
		return
	}
	if distanceToPlayer <= 5*gfx.ScreenWidth {
		b.UpdateFrequency = 60
		return
	}
	b.UpdateFrequency = 300

}

func (b *Brain) NeedsUpdating() bool {
	if b.owner.IsPlayer() {
		return false
	}
	if b.LastUpdated+int(b.UpdateFrequency) > theWorld.GetTick() {
		return false
	}
	return true
}

func (b *Brain) Init() {
	if b.TargetKey >= 0 && b.target == nil {
		b.target = theWorld.GetUnit(b.TargetKey)
	}
}

func (b *Brain) Idle() {
	if !randomizer.PercentageRoll(25) {
		return
	}
	x := int(b.owner.GetX())
	y := int(b.owner.GetY())
	reach := IdleMoveDistance * int(b.owner.GetF(Speed))
	newX := randomizer.RandomInt(x-reach, x+reach)
	newY := randomizer.RandomInt(y-reach, y+reach)
	b.owner.SetDestination(float64(newX), float64(newY))
}

func (b *Brain) ResolveState() {
	nearestEnemy := b.HasEnemyNearby()
	if nearestEnemy >= 0 {
		b.CurrentState = Chase
		b.TargetKey = nearestEnemy
		b.target = theWorld.GetUnit(nearestEnemy)
		return
	}
	b.CurrentState = Idle
}

func (b *Brain) Chase() {
	logger.General("Chasing "+theWorld.GetUnit(b.TargetKey).GetName(), nil)
	if !b.DistanceWithinVision(b.DistanceToUnit(b.target)) {
		logger.General("Lost target "+theWorld.GetUnit(b.TargetKey).GetName(), nil)
		b.target = nil
		b.TargetKey = -1
		b.CurrentState = Idle
		b.LastUpdated = 0
		return
	}
	b.owner.SetDestination(b.target.GetX(), b.target.GetY())
}

func (b *Brain) Attack() {

}

func (b *Brain) SetOwner(unit *Unit) {
	b.owner = unit
}

func (b *Brain) HasEnemyNearby() UnitKey {
	closestEnemy := UnitKey(-1)
	closestDistance := 999999
	for key, unit := range theWorld.GetUnits() {
		if key == b.owner.Id {
			continue
		}
		thisDistance := b.DistanceToUnit(unit)
		if !b.DistanceWithinVision(thisDistance) {
			continue
		}
		if thisDistance < closestDistance {
			closestDistance = thisDistance
		}
		closestEnemy = key
	}
	return closestEnemy
}

func (b *Brain) DistanceToUnit(u *Unit) int {
	return tiling.NewCoordF(b.owner.GetPos()).ChebyshevDist(tiling.NewCoordF(u.GetPos()))
}

func (b *Brain) DistanceWithinVision(distance int) bool {
	return distance < int(b.owner.GetF(Vision))
}
