package linkedlist

// Node represents each element in the linked list.
type Node struct {
	data int
	next *Node
}

// LinkedList structure.
type LinkedList struct {
	head *Node
	size int
}

// AddFirst inserts a node at the beginning.
func (l *LinkedList) AddFirst(value int) {
}

// AddLast inserts a node at the end.
func (l *LinkedList) AddLast(value int) {
	NewNode := &Node{data: value}

	if l.head == nil {
		l.head = NewNode
		l.size++
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = NewNode
	l.size++
}

// AddAt inserts at a specific index.
func (l *LinkedList) AddAt(index int, value int) {
	current := l.head
	prev := current
	for i := 0; i < index-1; i++ {
		prev = current
		current = current.next
	}
	newNode := Node{data: value, next: current}
	prev.next = &newNode
}

// RemoveFirst deletes the first node.
func (l *LinkedList) RemoveFirst() {
	if l.head.next == nil {
		l.head = nil
	}
	l.head = l.head.next
}

// RemoveLast deletes the last node.
func (l *LinkedList) RemoveLast() {
	currect := l.head

	for i := 0; i < l.size-1; i++ {
		currect = l.head.next
	}
	currect.next = nil
}

// RemoveAt deletes node at given index.
func (l *LinkedList) RemoveAt(index int) {
	currect := l.head
	for i := 0; i < index-1; i++ {
		currect = l.head.next
	}
	currect.next = currect.next.next
}

// Find searches for a value and returns its index.
func (l *LinkedList) Find(value int) int {
	current := l.head

	for i := 0; i < l.size; i++ {
		if current.data == value {
			return i
		}
		current = current.next
	}

	return 0
}

// Length returns size of list.
func (l *LinkedList) Length() int {
	return l.size
}

// Print displays the linked list.
func (l *LinkedList) Print() {
	currect := l.head

	for currect != nil {
		l.head = l.head.next
	}
}
