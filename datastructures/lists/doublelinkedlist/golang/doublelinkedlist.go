package doublelinkedlist

type DoubleLinkedNode struct {
	value      interface{}
	prev, next *DoubleLinkedNode
}

func NewDoubleLinkedNode(value interface{}, prev, next *DoubleLinkedNode) *DoubleLinkedNode {
	return &DoubleLinkedNode{
		value: value,
		prev:  prev,
		next:  next,
	}
}

func (n *DoubleLinkedNode) Prev() *DoubleLinkedNode {
	return n.prev
}

func (n *DoubleLinkedNode) Next() *DoubleLinkedNode {
	return n.next
}

func (n *DoubleLinkedNode) Value() interface{} {
	return n.value
}

type DoubleLinkedList struct {
	first, last *DoubleLinkedNode
	len         int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		first: nil,
		last:  nil,
		len:   0,
	}
}

func (l *DoubleLinkedList) First() *DoubleLinkedNode {
	return l.first
}

func (l *DoubleLinkedList) Last() *DoubleLinkedNode {
	return l.last
}

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

func (l *DoubleLinkedList) Len() int {
	return l.len
}

func (l *DoubleLinkedList) Clear() {
	l.first, l.last = nil, nil
	l.len = 0
}
