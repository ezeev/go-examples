package stack

//FILO
type Stack struct {
	items []int
}

func (s *Stack) Push(n int) {
	s.items = append(s.items, n)
}

func (s *Stack) Pop() int {
	n := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return n
}

func (s *Stack) Items() []int {
	return s.items
}

func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}
