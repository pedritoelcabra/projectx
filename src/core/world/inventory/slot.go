package inventory

type Slot struct {
	Type      int
	Count     int
	StackSize int
}

func NewSlot() *Slot {
	aSlot := &Slot{}
	aSlot.Type = 0
	aSlot.Count = 0
	aSlot.StackSize = 0
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
