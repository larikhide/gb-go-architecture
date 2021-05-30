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
	l.len++
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

}

func (l *List) Append(node *Node) {
	l.Add(l.tail, node)
}

func (l *List) Preppend(node *Node) {
	l.Add(nil, node)
}

func (l *List) Delete(node *Node) {
	// если список состоит из 1 го элемента
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	}

	if l.head != nil {
		for tmp := l.head; tmp != l.tail; tmp = tmp.next {
			if tmp.next == node && node != l.tail {
				tmp.next = node.next
			}

			if tmp.next == node && node == l.tail {
				tmp.next = nil
				l.tail = tmp
			}
		}
	}
	if node == l.head {
		l.head = node.next
	}
	l.len--
}
