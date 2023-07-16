package heap

import "sync"

type Heap[T any] struct {
	data       []T
	comparator func(T, T) bool
	lock       sync.RWMutex
}

// create an instance of heap with comparator
func Init[T any](f func(T, T) bool) *Heap[T] {
	return &Heap[T]{
		data:       make([]T, 0),
		comparator: f,
	}
}

// check if heap is empty
func (h *Heap[T]) Empty() bool {
	return len(h.data) == 0
}

// return the number of elements in heap
func (h *Heap[T]) Size() uint64 {
	return uint64(len(h.data))
}

// return top element of heap
func (h *Heap[T]) Top() T {
	h.lock.RLock()
	defer h.lock.RUnlock()

	if h.Empty() {
		panic("Heap has no element.")
	}

	return h.data[0]
}

// push an element into heap
func (h *Heap[T]) Push(item T) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.data = append(h.data, item)

	j := h.Size() - 1
	for {
		i := uint64((int(j) - 1) / 2) // parent
		if i == j || !h.comparator(h.data[i], h.data[j]) {
			break
		}
		h.data[j], h.data[i] = h.data[i], h.data[j]
		j = i
	}
}

// remove an element from top of heap
func (h *Heap[T]) Pop() {
	h.lock.Lock()
	defer h.lock.Unlock()

	if h.Empty() {
		panic("Heap has no element.")
	}

	n := h.Size() - 1
	h.data[0], h.data[n] = h.data[n], h.data[0]
	h.data = h.data[0:n]

	var i uint64 = 0
	if n == 0 {
		return
	}
	n = n - 1
	for {
		var left uint64 = 2*i + 1

		if left > n {
			break
		}

		t := left
		if right := left + 1; right <= n && !h.comparator(h.data[right], h.data[left]) {
			t = right
		}

		if h.comparator(h.data[t], h.data[i]) {
			break
		}

		h.data[t], h.data[i] = h.data[i], h.data[t]
		i = t
	}
}
