// list is a doubly linked list implemented by generic
package list

import "sync"

type List[T any] struct {
	root Node[T]
	len  uint64
	lock sync.RWMutex
}

type Node[T any] struct {
	prev  *Node[T]
	next  *Node[T]
	list  *List[T]
	Value T
}

// return the previous node of current node
func (n *Node[T]) Prev() *Node[T] {
	if n.list != nil {
		n.list.lock.Lock()
		defer n.list.lock.Unlock()
	}

	if n == nil || n.prev == &n.list.root {
		return nil
	}
	return n.prev
}

// return the next node of current node
func (n *Node[T]) Next() *Node[T] {
	if n.list != nil {
		n.list.lock.Lock()
		defer n.list.lock.Unlock()
	}

	if n == nil || n.next == &n.list.root {
		return nil
	}
	return n.next
}

// create an instance of list
func Init[T any]() *List[T] {
	list := new(List[T])
	list.root.prev = &list.root
	list.root.next = &list.root
	list.root.list = list
	list.len = 0
	return list
}

// return the number of nodes in list
func (l *List[T]) Len() uint64 {
	return l.len
}

// return the first node in list
func (l *List[T]) Front() *Node[T] {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if l.Len() == 0 {
		return nil
	}
	return l.root.next
}

// return the last node in list
func (l *List[T]) Back() *Node[T] {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if l.Len() == 0 {
		return nil
	}
	return l.root.prev
}

// insert a node after target node
func (l *List[T]) insertAfter(val T, target *Node[T]) *Node[T] {
	if target.list != l {
		return nil
	}

	newNode := &Node[T]{
		Value: val,
	}

	newNode.prev = target
	newNode.next = target.next
	target.next.prev = newNode
	target.next = newNode
	newNode.list = l
	l.len++

	return newNode
}

// push a node with specified value at the front of list
func (l *List[T]) PushFront(val T) *Node[T] {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.insertAfter(val, &l.root)
}

// push a node with specified value at the back of list
func (l *List[T]) PushBack(val T) *Node[T] {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.insertAfter(val, l.root.prev)
}

// insert a node with specified value before target node
func (l *List[T]) InsertBefore(val T, target *Node[T]) *Node[T] {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.insertAfter(val, target.prev)
}

// insert a node with specified value after target node
func (l *List[T]) InsertAfter(val T, target *Node[T]) *Node[T] {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.insertAfter(val, target)
}

// remove a node in list
func (l *List[T]) remove(target *Node[T]) {
	target.prev.next = target.next
	target.next.prev = target.prev
	target.prev = nil
	target.next = nil
	target.list = nil
	l.len--
}

// remove a node in list and return its value
// panic if target is nil or not in list
func (l *List[T]) Remove(target *Node[T]) T {
	l.lock.Lock()
	defer l.lock.Unlock()

	if target == nil || target.list != l {
		panic("remove invalid node")
	}
	l.remove(target)
	return target.Value
}
