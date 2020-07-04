package inventory

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"strconv"
)

const (
	DefaultSlotCount     = 20
	DefaultSlotStackSize = 50
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
		anInventory.Slots[i] = NewSlot(DefaultSlotStackSize)
	}
	return anInventory
}

func (i *Inventory) AddItem(itemType int, amount int) int {
	bestSlot := -1
	for k := 0; k < i.SlotCount; k++ {
		if i.Slots[k].IsFull() {
			continue
		}
		if !i.Slots[k].IsEmpty() && i.Slots[k].GetType() != itemType {
			continue
		}
		if bestSlot < 0 {
			bestSlot = k
		}
		if i.Slots[k].GetType() == itemType {
			bestSlot = k
			break
		}
	}
	if bestSlot < 0 {
		return amount
	}
	if i.Slots[bestSlot].EmptySpace() >= amount {
		i.Slots[bestSlot].AddItems(itemType, amount)
		return 0
	}
	remain := amount - i.Slots[bestSlot].EmptySpace()
	i.Slots[bestSlot].AddItems(itemType, i.Slots[bestSlot].EmptySpace())
	return remain
}

/**
Returns number of items removed
*/
func (i *Inventory) RemoveItems(itemType int, amount int) int {
	amountRemaining := amount
	for k := i.SlotCount - 1; k >= 0; k-- {
		if i.Slots[k].IsEmpty() {
			continue
		}
		if i.Slots[k].GetType() != itemType {
			continue
		}
		amountRemaining -= i.Slots[k].RemoveItems(amount)
		if amountRemaining == 0 {
			break
		}
	}
	return amount - amountRemaining
}

func (i *Inventory) GetContentList() string {
	list := ""
	emptySlots := 0

	for k := 0; k < i.SlotCount; k++ {
		if i.Slots[k].IsEmpty() {
			emptySlots++
			continue
		}
		itemDef := defs.GetMaterialDefByKey(i.Slots[k].Type)
		list += "\n" + strconv.Itoa(i.Slots[k].GetCount()) + " " + itemDef.Name
	}
	list += "\n" + strconv.Itoa(emptySlots) + " empty slots"
	return list
}
