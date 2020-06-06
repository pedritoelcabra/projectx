package world

import (
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	utils2 "github.com/pedritoelcabra/projectx/src/core/world/utils"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"log"
	"strconv"
)

type Brain struct {
	owner           *Unit
	Type            int
	CurrentState    int
	LastUpdated     int
	TargetKey       UnitKey
	UpdateFrequency int
	target          *Unit
	WorkTarget      BuildingPointer
}

func NewBrain() *Brain {
	aBrain := &Brain{}
	aBrain.CurrentState = StateIdle
	aBrain.TargetKey = 0
	aBrain.LastUpdated = 0
	aBrain.UpdateFrequency = 1
	return aBrain
}

const (
	IdleMoveDistance = 10
	IdleMoveChance   = 25
)
const (
	StateIdle = iota
	StateChase
	StateFlee
	StateAttack
	StatePatrol
	StateWork
	StateReturn
)

func (b *Brain) GetOccupationString() string {
	if b.CurrentState == StateIdle {
		return "Idling..."
	}
	if b.CurrentState == StateWork {
		return "Working hard"
	}
	if b.CurrentState == StateReturn {
		return "Going home"
	}
	description := ""
	if b.CurrentState == StateChase {
		description = "Chasing"
	}
	if b.CurrentState == StateAttack {
		description = "Attacking"
	}
	if b.target != nil {
		description += " " + b.target.GetName()
	}
	return description
}

func (b *Brain) ProcessState() {
	if b.owner.IsBusy() {
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
	case StateWork:
		b.Work()
		return
	case StateReturn:
		b.Return()
		return
	}
	log.Fatal("Unknown state: " + strconv.Itoa(b.CurrentState))
}

func (b *Brain) SetUpdateFrequency() {
	if b.CurrentState != StateIdle {
		b.UpdateFrequency = 10
	}
	distanceToPlayer := b.owner.DistanceToUnit(theWorld.PlayerUnit.unit)
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
	if !randomizer.PercentageRoll(10) {
		return
	}
	x := int(b.owner.GetX())
	y := int(b.owner.GetY())
	reach := IdleMoveDistance * int(b.owner.GetF(Speed))
	newX := randomizer.RandomInt(x-reach, x+reach)
	newY := randomizer.RandomInt(y-reach, y+reach)
	b.owner.SetDestination(float64(newX), float64(newY))
}

func (b *Brain) Return() {
	if b.owner.IsMoving() {
		return
	}
	home := b.owner.GetHome()
	if home == nil {
		return
	}
	if home.GetTile().GetCoord().Equals(b.owner.GetTileCoord()) {
		b.ResetState()
		return
	}
	homeSector := home.GetSector()
	if homeSector == nil {
		return
	}
	currentTile := b.owner.GetTile()
	currentSector := currentTile.GetSector()
	if currentSector == homeSector {
		b.ResetState()
		return
	}
	b.owner.SetDestination(home.GetTile().GetCenterPos())
}

func (b *Brain) Work() {
	if b.owner.Work == nil {
		b.ResetState()
		return
	}
	target := b.owner.Work.GetBuilding()
	if target == nil {
		b.ResetState()
		return
	}
	targetCoord := target.GetWorkLocation()
	targetTile := theWorld.Grid.Tile(targetCoord)
	if !targetTile.GetCoord().Equals(b.owner.GetTile().GetCoord()) {
		b.owner.SetDestination(targetTile.GetCenterPos())
		return
	}
	target.AddWork()
}

func (b *Brain) ForceUpdate() {
	b.LastUpdated = 0
}

func (b *Brain) ResetState() {
	b.target = nil
	b.TargetKey = -1
	b.CurrentState = StateIdle
	b.LastUpdated = 0
	b.ResolveState()
}

func (b *Brain) ResolveState() {
	nearestEnemy := b.owner.ClosestVisibleEnemy()
	if nearestEnemy >= 0 {
		b.CurrentState = StateChase
		b.TargetKey = nearestEnemy
		b.target = theWorld.GetUnit(nearestEnemy)
		return
	}
	if !b.owner.IsInOwnedSector() {
		b.CurrentState = StateReturn
		return
	}
	if b.owner.CanWork() {
		sector := b.owner.GetTile().GetSector()
		if sector != nil {
			job := sector.GetEmptyWorkPlace()
			if job != nil {
				job.HireWorker(b.owner)
				b.CurrentState = StateWork
				return
			}
		}
	}
	b.CurrentState = StateIdle
}

func (b *Brain) Chase() {
	if !b.target.IsAlive() {
		b.ResetState()
		return
	}
	distance := b.owner.DistanceToUnit(b.target)
	//logger.General("Chasing "+theWorld.GetUnit(b.TargetKey).GetName()+" distance "+strconv.Itoa(distance), nil)
	if !b.owner.DistanceWithinVision(distance) {
		//logger.General("Lost target "+theWorld.GetUnit(b.TargetKey).GetName(), nil)
		b.ResetState()
		return
	}
	if b.owner.DistanceWithinAttackRange(distance) {
		b.CurrentState = StateAttack
		b.ForceUpdate()
		return
	}
	b.owner.SetDestination(b.PositionToAttackTarget())
}

func (b *Brain) PositionToAttackTarget() (x, y float64) {
	return utils2.AdvanceAlongLine(b.target.GetX(), b.target.GetY(), b.owner.GetX(), b.owner.GetY(), b.owner.GetF(AttackRange)-2.0)
}

func (b *Brain) Attack() {
	if !b.target.IsAlive() {
		b.ResetState()
		return
	}
	distance := b.owner.DistanceToUnit(b.target)
	if !b.owner.DistanceWithinAttackRange(distance) {
		b.CurrentState = StateChase
		b.ForceUpdate()
		return
	}
	b.PerformAttackOn(b.target)
}

func (b *Brain) PerformAttackOn(target *Unit) {
	b.LastUpdated = theWorld.GetTick()
	b.owner.PerformAttackOn(target)
}

func (b *Brain) SetOwner(unit *Unit) {
	b.owner = unit
}
