package stack

import "sync"

type Stack[T any] struct {
	data []T
	top  uint64
	lock sync.RWMutex
}

// create an instance of stack with specified capacity
func Init[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

// check if stack is empty
func (s *Stack[T]) Empty() bool {
	return s.top == 0
}

// return the number of elements in stack
func (s *Stack[T]) Size() uint64 {
	return s.top
}

// return top element of stack
func (s *Stack[T]) Top() T {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if s.Empty() {
		panic("Stack has no element.")
	}

	return s.data[s.top-1]
}

// push an element into staRWMutex
func (s *Stack[T]) Push(item T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = append(s.data, item)
	s.top++
}

// remove an element from top of stack
func (s *Stack[T]) Pop() {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.Empty() {
		panic("Stack has no element.")
	}

	s.data = s.data[:s.Size()-1]
	s.top--
}
