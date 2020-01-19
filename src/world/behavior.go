package world

import (
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"log"
	"strconv"
)

type Brain struct {
	owner        *Unit
	Type         int
	CurrentState int
	LastUpdated  int
	Target       UnitKey
}

func NewBrain() *Brain {
	aBrain := &Brain{}
	aBrain.CurrentState = Idle
	aBrain.Target = 0
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
	switch b.CurrentState {
	case Idle:
		b.Idle()
		return
	case Chase:
		b.Chase()
		return
	case Flee:
		b.Flee()
		return
	case Attack:
		b.Attack()
		return
	case Patrol:
		b.Patrol()
		return
	}
	log.Fatal("Unknown state: " + strconv.Itoa(b.CurrentState))
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

func (b *Brain) Chase() {

}

func (b *Brain) Flee() {

}

func (b *Brain) Attack() {

}

func (b *Brain) Patrol() {

}

func (b *Brain) SetOwner(unit *Unit) {
	b.owner = unit
}
