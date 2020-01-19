package world

import (
	"github.com/pedritoelcabra/projectx/src/core/logger"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"log"
	"strconv"
)

type Brain struct {
	owner        *Unit
	Type         int
	CurrentState int
	LastUpdated  int
	TargetKey    UnitKey
	target       *Unit
}

func NewBrain() *Brain {
	aBrain := &Brain{}
	aBrain.CurrentState = Idle
	aBrain.TargetKey = 0
	aBrain.LastUpdated = 0
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
	if b.owner.IsBusy() {
		return
	}
	if b.owner.GetF(UpdateFrequency) == 0.0 {
		return
	}
	if b.LastUpdated+int(b.owner.GetF(UpdateFrequency)) > theWorld.GetTick() {
		return
	}
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
