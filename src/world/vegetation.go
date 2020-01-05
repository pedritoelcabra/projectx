package world

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/gfx"
)

type Vegetation struct {
	Name      string
	Template  *defs.VegetationDef
	Sprite    gfx.Sprite `json:"-"`
	SpriteKey gfx.SpriteKey
}

func NewVegetation(name string) *Vegetation {
	vegDefs := defs.VegetationDefs()
	aVegetation := &Vegetation{}
	aVegetation.Template = vegDefs[name]
	aVegetation.SpriteKey = gfx.GetSpriteKey(aVegetation.Template.GetGraphic())
	aVegetation.Name = name
	aVegetation.Init()
	return aVegetation
}

func (v *Vegetation) GetName() string {
	return v.Name
}

func (v *Vegetation) GetMovementCost() float64 {
	return v.Template.MovementCost
}

func (v *Vegetation) Init() {
	v.Sprite = gfx.NewHexSprite(v.SpriteKey)
}