package cypher_go_dsl

import (
"fmt"
)

func main() {
	// Create a new Deque
	deque := NewDeque()

	// Inject two items in back
	deque.Inject("Second")
	deque.Inject("Third")
	// Push an item in front
	deque.Push("First")

	fmt.Println(deque.Items)

	// Remove an item in front
	deque.Pop()

	// Remove an item in back
	deque.Eject()

	// Check if the deque is empty and his values
	fmt.Println(deque.IsEmpty(), deque.Items)
}

func NewDeque() *Deque {
	return &Deque{}
}

type Deque struct {
	Items []interface{}
}

func (s *Deque) Push(item interface{}) {
	temp := []interface{}{item}
	s.Items = append(temp, s.Items...)
}

func (s *Deque) Inject(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Deque) Pop() interface{} {
	defer func() {
		s.Items = s.Items[1:]
	}()
	return s.Items[0]
}

func (s *Deque) Eject() interface{} {
	i := len(s.Items) - 1
	defer func() {
		s.Items = append(s.Items[:i], s.Items[i+1:]...)
	}()
	return s.Items[i]
}

func (s *Deque) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}
