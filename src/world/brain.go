package world

import (
	"github.com/pedritoelcabra/projectx/src/core/logger"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"github.com/pedritoelcabra/projectx/src/world/utils"
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
	aBrain.CurrentState = StateIdle
	aBrain.TargetKey = 0
	aBrain.LastUpdated = 0
	aBrain.BusyTime = 0
	aBrain.UpdateFrequency = 1
	return aBrain
}

const (
	StateIdle = iota
	StateChase
	StateFlee
	StateAttack
	StatePatrol

	IdleMoveDistance = 10
	IdleMoveChance   = 25
)

func (b *Brain) ProcessState() {
	if b.BusyTime > 0 {
		b.BusyTime--
		return
	}
	if !b.NeedsUpdating() {
		return
	}
	b.SetUpdateFrequency()
	//logger.General("Updating unit "+strconv.Itoa(int(b.owner.Id))+" on tick "+strconv.Itoa(theWorld.GetTick()), nil)
	b.LastUpdated = theWorld.GetTick()
	switch b.CurrentState {
	case StateIdle:
		b.Idle()
		return
	case StateChase:
		b.Chase()
		return
	case StateAttack:
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
	b.ResolveState()
	if b.CurrentState != StateIdle {
		b.ForceUpdate()
		return
	}
	if !randomizer.PercentageRoll(25) {
		return
	}
	b.BusyTime = 100
	x := int(b.owner.GetX())
	y := int(b.owner.GetY())
	reach := IdleMoveDistance * int(b.owner.GetF(Speed))
	newX := randomizer.RandomInt(x-reach, x+reach)
	newY := randomizer.RandomInt(y-reach, y+reach)
	b.owner.SetDestination(float64(newX), float64(newY))
}

func (b *Brain) ForceUpdate() {
	b.LastUpdated = 0
	b.BusyTime = 0
}

func (b *Brain) ResetState() {
	b.target = nil
	b.TargetKey = -1
	b.CurrentState = StateIdle
	b.LastUpdated = 0
	b.ResolveState()
}

func (b *Brain) ResolveState() {
	nearestEnemy := b.HasEnemyNearby()
	if nearestEnemy >= 0 {
		b.CurrentState = StateChase
		b.TargetKey = nearestEnemy
		b.target = theWorld.GetUnit(nearestEnemy)
		return
	}
	b.CurrentState = StateIdle
}

func (b *Brain) Chase() {
	distance := b.DistanceToUnit(b.target)
	logger.General("Chasing "+theWorld.GetUnit(b.TargetKey).GetName()+" distance "+strconv.Itoa(distance), nil)
	if !b.DistanceWithinVision(distance) {
		logger.General("Lost target "+theWorld.GetUnit(b.TargetKey).GetName(), nil)
		b.ResetState()
		return
	}
	if b.DistanceWithinAttackRange(distance) {
		b.CurrentState = StateAttack
		b.ForceUpdate()
		return
	}
	b.owner.SetDestination(b.PositionToAttackTarget())
}

func (b *Brain) PositionToAttackTarget() (x, y float64) {
	return utils.AdvanceAlongLine(b.target.GetX(), b.target.GetY(), b.owner.GetX(), b.owner.GetY(), b.owner.GetF(AttackRange)-2.0)
}

func (b *Brain) Attack() {
	distance := b.DistanceToUnit(b.target)
	if !b.DistanceWithinAttackRange(distance) {
		b.CurrentState = StateChase
		b.ForceUpdate()
		return
	}
	b.PerformAttackOn(b.target)
}

func (b *Brain) PerformAttackOn(target *Unit) {
	attackSpeed := int(6000 / b.owner.GetF(AttackSpeed))
	b.BusyTime = attackSpeed
	b.LastUpdated = 0
	b.owner.PerformAttackOn(target)
	x, y := target.GetPos()
	b.owner.QueueAttackAnimation(x, y, attackSpeed)
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
		if !b.owner.GetFaction().IsHostileTowards(unit.GetFaction()) {
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

func (b *Brain) DistanceWithinAttackRange(distance int) bool {
	attackRange := int(b.owner.GetF(AttackRange))
	return distance <= attackRange
}
