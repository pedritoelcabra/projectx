package inventory

import (
	"github.com/pedritoelcabra/projectx/src/core/logger"
	"strconv"
)

const (
	DefaultSlotCount = 50
)

type SlotList map[int]*Slot

type Inventory struct {
	Slots     SlotList
	SlotCount int
}

func NewInventory(slotCount int) *Inventory {
	anInventory := &Inventory{}
	anInventory.SlotCount = slotCount
	anInventory.Slots = make(SlotList)
	for i := 0; i < slotCount; i++ {
		anInventory.Slots[i] = NewSlot()
	}
	return anInventory
}

func (i *Inventory) AddItem(itemType int, amount int) {
	logger.General("added "+strconv.Itoa(amount)+" x "+strconv.Itoa(itemType), nil)
}
