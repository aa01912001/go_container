package queue

import "sync"

type Queue[T any] struct {
	data []T
	lock sync.RWMutex
}

// create an instance of queue
func Init[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

// check if queue is empty
func (q *Queue[T]) Empty() bool {
	return len(q.data) == 0
}

// return the number of elements in queue
func (q *Queue[T]) Size() uint64 {
	return uint64(len(q.data))
}

// return front element of queue
func (q *Queue[T]) Front() T {
	q.lock.RLock()
	defer q.lock.RUnlock()

	if q.Empty() {
		panic("Queue has no element.")
	}

	return q.data[0]
}

// push an element into queue
func (q *Queue[T]) Push(item T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.data = append(q.data, item)
}

// remove an element from front of stack
func (q *Queue[T]) Pop() {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.Empty() {
		panic("Queue has no element.")
	}

	q.data = q.data[1:]
}
