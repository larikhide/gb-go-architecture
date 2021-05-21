package models

type LinkedList interface {
	Head() *Node
	Tail() *Node
	Len() *Node

	Add(prev *Node, node *Node)
	Append(node *Node)
	Preppend(node *Node)
	Delete(node *Node)
}

type Node struct {
	Data int
	prev *Node
	next *Node
}

type List struct {
	len  int
	head *Node
	tail *Node
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Tail() *Node {
	return l.tail
}

func (l *List) Len() int {
	return l.len
}

// Add insert new node *Node previous prev *Node
func (l *List) Add(prev *Node, node *Node) {

	//если хотим вставить ноду в самое начало списка
	if prev == nil {
		node.next = l.head
		l.head = node
		return
	}
	// если в листе нет ни одной ноды
	if l.head == nil {
		l.head = node
		l.tail = l.head
		return
	}

	node.next = prev
	node.prev = prev.prev
	prev.prev = node

	l.len++
}

func (l *List) Append(node *Node) {
	return l.Add(l.tail, node)
}

func (l *List) Preppend(node *Node) {
	return l.Add(nil, node)
}

func (l *List) Delete(node *Node) {

	//TODO: описать удаление определенной ноды из списка

	l.len--
}
