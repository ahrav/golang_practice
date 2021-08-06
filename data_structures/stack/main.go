package main

type Stack struct {
	items []int
}

// Push will add a value to the end.
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove value at the end and return it.
func (s *Stack) Pop() int{
	l := len(s.items)
	val := s.items[l]
	s.items = s.items[:l]
	return val
}

func main() {
	myStack := Stack{}
	myStack.Push(2)
	myStack.Push(3)
}
