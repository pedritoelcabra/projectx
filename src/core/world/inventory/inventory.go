package inventory

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
