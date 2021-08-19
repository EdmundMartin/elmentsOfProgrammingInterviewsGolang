package stacks_queues

type ElementWithCachedMax struct {
	Element int
	Max     int
}

type StackWithMax struct {
	Elements []*ElementWithCachedMax
}

func NewStackWithMax() *StackWithMax {
	return &StackWithMax{Elements: []*ElementWithCachedMax{}}
}

func (s *StackWithMax) Empty() bool {
	return len(s.Elements) == 0
}

func (s *StackWithMax) Max() int {
	return s.Elements[len(s.Elements)-1].Max
}

func (s *StackWithMax) Pop() int {
	popped, elements := s.Elements[len(s.Elements)-1], s.Elements[:len(s.Elements)-1]
	s.Elements = elements
	return popped.Element
}

func (s *StackWithMax) getMax(newValue int) int {
	if s.Empty() {
		return newValue
	}
	currentMax := s.Max()
	if newValue > currentMax {
		return newValue
	}
	return currentMax
}

func (s *StackWithMax) Push(value int) {
	s.Elements = append(s.Elements, &ElementWithCachedMax{
		Element: value,
		Max:     s.getMax(value),
	})
}
