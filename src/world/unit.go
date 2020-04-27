package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"github.com/pedritoelcabra/projectx/src/world/utils"
	"log"
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
	Template   *defs.UnitDef
}

func NewUnit(templateName string, location tiling.Coord) *Unit {
	template := defs.UnitDefs()[templateName]
	if template == nil {
		log.Fatal("Invalid Unit Template: " + templateName)
	}
	aUnit := &Unit{}
	aUnit.Template = template
	aUnit.Alive = true
	aUnit.Name = template.Name
	aUnit.X = float64(location.X())
	aUnit.Y = float64(location.Y())
	aUnit.Attributes = NewAttributes(template.Attributes)
	aUnit.SetToMaxHealth()
	aUnit.SetF(BusyTime, 0.0)
	aUnit.SetF(LastCombatAction, 0.0)
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
	u.PassiveHeal()
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

func (u *Unit) GetMovementSpeed() float64 {
	return u.GetF(Speed)
}

func (u *Unit) GetId() UnitKey {
	return u.Id
}

func (u *Unit) GetDescription() string {
	return u.Template.Description
}

func (u *Unit) GetStats() string {
	stats := u.Brain.GetOccupationString()
	stats += "\nHealth: " + gfx.HealthString(u)
	stats += "\nDamage: " + utils.NumberFormat(u.GetF(AttackDamage))
	stats += "\nAttack Speed: " + utils.NumberFormat(60/u.GetAttackCoolDown())
	stats += "\nMovement Speed: " + utils.NumberFormat(u.GetMovementSpeed())
	return stats
}
