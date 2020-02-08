package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/logger"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"github.com/pedritoelcabra/projectx/src/world/utils"
	"log"
	"math"
)

type UnitKey int
type UnitMap map[UnitKey]*Unit
type UnitArray []*Unit

type Unit struct {
	Id         UnitKey
	Sprite     gfx.Sprite `json:"-"`
	Unit       *ebiten.Image
	spriteKey  gfx.SpriteKey
	Graphics   map[string]string
	X          float64
	Y          float64
	DestX      float64
	DestY      float64
	Moving     bool
	Size       float64
	Name       string
	Attributes *Attributes
	Brain      *Brain
	Alive      bool
}

func NewUnit(templateName string, location tiling.Coord) *Unit {
	template := defs.UnitDefs()[templateName]
	if template == nil {
		log.Fatal("Invalid Unit Template: " + templateName)
	}
	aUnit := &Unit{}
	aUnit.Alive = true
	aUnit.Name = template.Name
	aUnit.X = float64(location.X())
	aUnit.Y = float64(location.Y())
	aUnit.Attributes = NewAttributes(template.Attributes)
	aUnit.SetToMaxHealth()
	aUnit.SetF(BusyTime, 0.0)
	aUnit.SetEquipmentGraphics(template)
	aUnit.Brain = NewBrain()
	aUnit.Init()
	aUnit.Id = theWorld.AddUnit(aUnit)
	if template.Faction != "" {
		aUnit.SetFaction(theWorld.GetFactionByName(template.Faction))
	}
	return aUnit
}

func (u *Unit) IsPlayer() bool {
	return u.Id == 0
}

func (u *Unit) IsAlive() bool {
	return u.Alive
}

func (u *Unit) IsBusy() bool {
	return u.GetF(BusyTime) > 0.0
}

func (u *Unit) DrawSprite(screen *gfx.Screen) {
	u.Sprite.DrawSprite(screen, u.X, u.Y)
	gfx.DrawHealthBar(u, screen)
}

func (u *Unit) ShouldDraw() bool {
	if !u.IsAlive() {
		return false
	}
	return EntityShouldDraw(u.GetX(), u.GetY())
}

func (u *Unit) SetPosition(x, y float64) {
	u.X = x
	u.Y = y
	u.CheckIfMoving()
}

func (u *Unit) SetDestination(x, y float64) {
	u.DestX = x
	u.DestY = y
	u.CheckIfMoving()
	if u.Moving {
		u.OrientateTowardsPoint(u.DestX, u.DestY)
	}
}

func (u *Unit) OrientateTowardsPoint(x, y float64) {
	if math.Abs(u.X-x)+1 > math.Abs(u.Y-y) {
		if u.X > x {
			u.Sprite.SetFacing(gfx.FaceLeft)
			return
		}
		u.Sprite.SetFacing(gfx.FaceRight)
		return
	}
	if u.Y > y {
		u.Sprite.SetFacing(gfx.FaceUp)
		return
	}
	u.Sprite.SetFacing(gfx.FaceDown)
}

func (u *Unit) CollidesWith(x, y float64) bool {
	return utils.CalculateDistance(u.X, u.Y, x, y) < u.GetF(Size)
}

func (u *Unit) Init() {
	u.SetGraphics()
	u.Brain.SetOwner(u)
	u.Brain.Init()
}

func (u *Unit) SetGraphics() {
	var spriteComposite []gfx.SpriteKey
	for _, slotName := range gfx.LpcCompositeSlotOrder() {
		if slotValue, ok := u.Graphics[slotName]; ok {
			spriteComposite = append(spriteComposite, gfx.GetLpcKey(slotValue))
		}
	}
	u.spriteKey = gfx.GetLpcComposite(spriteComposite)
	u.Sprite = gfx.NewLpcSprite(u.spriteKey)
}

func (u *Unit) SetEquipmentGraphics(unitDefinition *defs.UnitDef) {
	u.Graphics = make(map[string]string)
	for _, def := range unitDefinition.Equipments {
		if _, ok := u.Graphics[def.Slot]; ok {
			continue
		}
		u.Graphics[def.Slot] = defs.ResolveGraphicChance(unitDefinition.Equipments, def.Slot)
	}
}

func (u *Unit) Update() {
	if !u.IsAlive() {
		return
	}
	u.SetF(BusyTime, u.GetF(BusyTime)-1.0)
	u.Brain.ProcessState()
	if u.Moving {
		oldCoord := u.GetTileCoord()
		oldTile := theWorld.Grid.Tile(oldCoord)
		movementCost := oldTile.GetF(MovementCost)
		if movementCost == 0 {
			movementCost = 1.0
		}
		movementSpeed := u.GetF(Speed) / movementCost
		newX, newY := utils.AdvanceAlongLine(u.X, u.Y, u.DestX, u.DestY, movementSpeed)
		newCoord := tiling.PixelFToTileC(newX, newY)
		canMove := true
		if oldCoord != newCoord {
			newTile := theWorld.Grid.Tile(newCoord)
			if newTile.IsImpassable() {
				canMove = false
				u.StopMovement()
			}
		}
		if canMove {
			u.SetPosition(newX, newY)
		}
	}
}

func (u *Unit) CheckIfMoving() {
	if u.DestY != u.Y || u.DestX != u.X {
		u.Moving = true
		return
	}
	u.Moving = false
}

func (u *Unit) StopMovement() {
	u.DestY = u.X
	u.DestY = u.Y
	u.Moving = false
}

func (u *Unit) GetPos() (x, y float64) {
	return u.X, u.Y
}

func (u *Unit) GetTileCoord() tiling.Coord {
	return tiling.PixelFToTileC(u.GetPos())
}

func (u *Unit) QueueAttackAnimation(x, y float64, speed int) {
	u.Sprite.QueueAttackAnimation((x-u.GetX())/2, (y-u.GetY())/2, speed)
	u.OrientateTowardsPoint(x, y)
}

func (u *Unit) GetAttackCoolDown() float64 {
	return 6000 / u.GetF(AttackSpeed)
}

func (u *Unit) PerformAttackOn(target *Unit) {
	if u.IsBusy() {
		return
	}
	attackSpeed := u.GetAttackCoolDown()
	u.SetF(BusyTime, attackSpeed)
	x, y := target.GetPos()
	u.QueueAttackAnimation(x, y, int(attackSpeed))
	u.StopMovement()
	//logger.General("Attacking "+target.GetName(), nil)
	attack := NewAttack()
	attack.Damage = u.GetF(AttackDamage)
	attack.Attacker = u
	attack.Defender = target
	target.ReceiveAttack(attack)
}

func (u *Unit) ReceiveAttack(attack *Attack) {
	u.Attributes.ApplyF(HitPoints, -attack.Damage)
	if u.GetHealth() <= 0 {
		u.Alive = false
		logger.General(u.GetName()+" died", nil)
	}
}

func (u *Unit) SetToMaxHealth() {
	u.SetF(HitPoints, u.GetF(MaxHitPoints))
}

func (u *Unit) GetX() float64 {
	return u.X
}

func (u *Unit) GetY() float64 {
	return u.Y
}

func (u *Unit) GetName() string {
	return u.Name
}

func (u *Unit) GetFaction() *Faction {
	return theWorld.GetFaction(FactionKey(u.Get(FactionId)))
}

func (u *Unit) SetFaction(faction *Faction) {
	u.Set(FactionId, int(faction.GetId()))
}

func (u *Unit) Get(key int) int {
	return u.Attributes.Get(int(key))
}

func (u *Unit) GetF(key int) float64 {
	return u.Attributes.GetF(key)
}

func (u *Unit) Set(key, value int) {
	u.Attributes.Set(key, value)
}

func (u *Unit) SetF(key int, value float64) {
	u.Attributes.SetF(key, value)
}

func (u *Unit) GetHealth() float64 {
	return u.GetF(HitPoints)
}

func (u *Unit) GetMaxHealth() float64 {
	return u.GetF(MaxHitPoints)
}

func (u *Unit) ClosestVisibleEnemy() UnitKey {
	closestEnemy := UnitKey(-1)
	closestDistance := 999999
	for key, unit := range theWorld.GetUnits() {
		if !unit.IsAlive() {
			continue
		}
		if key == u.Id {
			continue
		}
		thisDistance := u.DistanceToUnit(unit)
		if !u.DistanceWithinVision(thisDistance) {
			continue
		}
		if !u.GetFaction().IsHostileTowards(unit.GetFaction()) {
			continue
		}
		if thisDistance < closestDistance {
			closestDistance = thisDistance
			closestEnemy = key
		}
	}
	return closestEnemy
}

func (u *Unit) DistanceToUnit(t *Unit) int {
	return tiling.NewCoordF(u.GetPos()).ChebyshevDist(tiling.NewCoordF(t.GetPos()))
}

func (u *Unit) DistanceWithinVision(distance int) bool {
	return distance < int(u.GetF(Vision))
}

func (u *Unit) DistanceWithinAttackRange(distance int) bool {
	attackRange := int(u.GetF(AttackRange))
	return distance <= attackRange
}
