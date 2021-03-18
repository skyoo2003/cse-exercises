// nolint
package doublelinkedlist

// DoubleLinkedNode double linked list's node data structure
type DoubleLinkedNode struct {
	value      interface{}
	prev, next *DoubleLinkedNode
}

// NewDoubleLinkedNode create a node
func NewDoubleLinkedNode(value interface{}, prev, next *DoubleLinkedNode) *DoubleLinkedNode {
	return &DoubleLinkedNode{
		value: value,
		prev:  prev,
		next:  next,
	}
}

// Prev get a previous node
func (n *DoubleLinkedNode) Prev() *DoubleLinkedNode {
	return n.prev
}

// Next get a next node
func (n *DoubleLinkedNode) Next() *DoubleLinkedNode {
	return n.next
}

// Value get a value of node
func (n *DoubleLinkedNode) Value() interface{} {
	return n.value
}

// DoubleLinkedList double linked list data structure
type DoubleLinkedList struct {
	first, last *DoubleLinkedNode
	len         int
}

// NewDoubleLinkedList create a double linked list
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		first: nil,
		last:  nil,
		len:   0,
	}
}

// First get a first node in double linked list
func (l *DoubleLinkedList) First() *DoubleLinkedNode {
	return l.first
}

// Last get a last node in double linked list
func (l *DoubleLinkedList) Last() *DoubleLinkedNode {
	return l.last
}

// InsertAfter insert a value after the node
func (l *DoubleLinkedList) InsertAfter(value interface{}, node *DoubleLinkedNode) *DoubleLinkedNode {
	newNode := NewDoubleLinkedNode(value, node, nil)
	if node.next == nil {
		l.last = newNode
	} else {
		newNode.next = node.next
		node.next.prev = newNode
	}
	node.next = newNode
	l.len++
	return newNode
}

// InsertBefore insert a value before the node
func (l *DoubleLinkedList) InsertBefore(value interface{}, node *DoubleLinkedNode) *DoubleLinkedNode {
	newNode := NewDoubleLinkedNode(value, nil, node)
	if node.prev == nil {
		l.first = newNode
	} else {
		newNode.prev = node.prev
		node.prev.next = newNode
	}
	node.prev = newNode
	l.len++
	return newNode
}

// InsertFront insert a value in front of double linked list
func (l *DoubleLinkedList) InsertFront(value interface{}) *DoubleLinkedNode {
	var newNode *DoubleLinkedNode
	if l.first == nil {
		newNode = NewDoubleLinkedNode(value, nil, nil)
		l.first, l.last = newNode, newNode
	} else {
		newNode = l.InsertBefore(value, l.first)
	}
	l.len++
	return newNode
}

// InsertLast insert a value in the end of double linked list
func (l *DoubleLinkedList) InsertLast(value interface{}) *DoubleLinkedNode {
	var newNode *DoubleLinkedNode
	if l.last == nil {
		newNode = l.InsertFront(value)
	} else {
		newNode = l.InsertAfter(value, l.last)
	}
	l.len++
	return newNode
}

// Remove delete a node from double linked list
func (l *DoubleLinkedList) Remove(node *DoubleLinkedNode) interface{} {
	value := node.value
	if node.prev == nil {
		l.first = node.next
	} else {
		node.prev.next = node.next
	}
	if node.next == nil {
		l.last = node.prev
	} else {
		node.next.prev = node.prev
	}
	l.len--
	return value
}

// Len get size of double linked list
func (l *DoubleLinkedList) Len() int {
	return l.len
}

// Clear reset the double linked list
func (l *DoubleLinkedList) Clear() {
	l.first, l.last = nil, nil
	l.len = 0
}
