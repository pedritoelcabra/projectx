package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/container"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"github.com/pedritoelcabra/projectx/src/world/utils"
	"log"
	"math"
)

type UnitKey int
type UnitMap map[UnitKey]*Unit

type Unit struct {
	Id        UnitKey
	Sprite    gfx.Sprite `json:"-"`
	Unit      *ebiten.Image
	spriteKey gfx.SpriteKey
	Graphics  map[string]string
	X         float64
	Y         float64
	DestX     float64
	DestY     float64
	Moving    bool
	Speed     float64
	Data      *container.Container
	Size      float64
	Name      string
}

func NewUnit(templateName string, location tiling.Coord) *Unit {
	template := defs.UnitDefs()[templateName]
	if template == nil {
		log.Fatal("Invalid Unit Template: " + templateName)
	}
	aUnit := &Unit{}
	aUnit.Name = template.Name
	aUnit.X = float64(location.X())
	aUnit.Y = float64(location.Y())
	aUnit.SetEquipmentGraphics(template)
	aUnit.Init()
	aUnit.Speed = 100
	aUnit.Size = float64(gfx.DefaultCollisionSize)
	aUnit.Data = container.NewContainer()
	aUnit.Id = theWorld.AddUnit(aUnit)
	return aUnit
}

func (u *Unit) DrawSprite(screen *gfx.Screen) {
	u.Sprite.DrawSprite(screen, u.X, u.Y)
}

func (u *Unit) ShouldDraw() bool {
	return EntityShouldDraw(u.GetX(), u.GetY())
}

func (u *Unit) SetPosition(x, y float64) {
	u.X = x
	u.Y = y
	u.DestX = x
	u.DestY = y
	u.CheckIfMoving()
}

func (u *Unit) SetDestination(x, y float64) {
	u.DestX = x
	u.DestY = y
	u.CheckIfMoving()
	u.CheckOrientation()
}

func (u *Unit) CheckOrientation() {
	if !u.Moving {
		return
	}
	if math.Abs(u.X-u.DestX)+1 > math.Abs(u.Y-u.DestY) {
		if u.X > u.DestX {
			u.Sprite.SetFacing(gfx.FaceLeft)
			return
		}
		u.Sprite.SetFacing(gfx.FaceRight)
		return
	}
	if u.Y > u.DestY {
		u.Sprite.SetFacing(gfx.FaceUp)
		return
	}
	u.Sprite.SetFacing(gfx.FaceDown)
}

func (u *Unit) SetSpeed(speed float64) {
	u.Speed = speed
}

func (u *Unit) SetSize(size float64) {
	u.Size = size
}

func (u *Unit) CollidesWith(x, y float64) bool {
	return utils.CalculateDistance(u.X, u.Y, x, y) < u.Size
}

func (u *Unit) Init() {
	u.SetGraphics()
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

func (u *Unit) Update(tick int, grid *Grid) {
	if u.Moving {
		oldCoord := tiling.PixelFToTileC(u.GetPos())
		oldTile := grid.Tile(oldCoord)
		movementCost := oldTile.GetF(MovementCost)
		if movementCost == 0 {
			movementCost = 1.0
		}
		movementSpeed := u.Speed / movementCost
		newX, newY := utils.AdvanceAlongLine(u.X, u.Y, u.DestX, u.DestY, movementSpeed)
		newCoord := tiling.PixelFToTileC(newX, newY)
		canMove := true
		if oldCoord != newCoord {
			newTile := grid.Tile(newCoord)
			if newTile.IsImpassable() {
				canMove = false
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

func (u *Unit) GetPos() (x, y float64) {
	return u.X, u.Y
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

func (u *Unit) Get(key int) int {
	return u.Data.Get(key)
}

func (u *Unit) GetF(key int) float64 {
	return u.Data.GetF(key)
}

func (u *Unit) Set(key, value int) {
	u.Data.Set(key, value)
}

func (u *Unit) SetF(key int, value float64) {
	u.Data.SetF(key, value)
}
