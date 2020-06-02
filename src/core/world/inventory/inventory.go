package inventory

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"strconv"
)

const (
	DefaultSlotCount     = 50
	DefaultSlotStackSize = 100
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

func (i *Inventory) GetContentList() string {
	list := ""
	items := make(map[int]int)

	for k := 0; k < i.SlotCount; k++ {
		if i.Slots[k].IsEmpty() {
			continue
		}
		if _, ok := items[i.Slots[k].Type]; !ok {
			items[i.Slots[k].Type] = 0
		}
		items[i.Slots[k].Type] += i.Slots[k].GetCount()
	}

	for key, count := range items {
		itemDef := defs.GetMaterialDefByKey(key)
		list += "\n" + strconv.Itoa(count) + " " + itemDef.Name
	}
	return list
}
