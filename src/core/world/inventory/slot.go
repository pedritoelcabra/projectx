package inventory

type Slot struct {
	Type      int
	Count     int
	StackSize int
}

func NewSlot(stackSize int) *Slot {
	aSlot := &Slot{}
	aSlot.Type = 0
	aSlot.Count = 0
	aSlot.StackSize = stackSize
	return aSlot
}

func (s *Slot) IsEmpty() bool {
	return s.Count == 0
}

func (s *Slot) IsFull() bool {
	return s.Count == s.StackSize
}

func (s *Slot) EmptySpace() int {
	return s.StackSize - s.Count
}

func (s *Slot) GetCount() int {
	return s.Count
}

func (s *Slot) SetCount(count int) {
	s.Count = count
}

func (s *Slot) GetType() int {
	return s.Type
}

func (s *Slot) SetType(itemType int) {
	s.Type = itemType
}

func (s *Slot) AddItems(itemType int, amount int) bool {
	if s.Type != itemType {
		if !s.IsEmpty() {
			return false
		}
		s.SetType(itemType)
	}
	if s.EmptySpace() < amount {
		return false
	}
	s.Count += amount
	return true
}

func (s *Slot) RemoveItems(amount int) bool {
	if s.Count < amount {
		return false
	}
	s.Count -= amount
	return true
}
