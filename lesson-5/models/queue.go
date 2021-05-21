package models

type QueueInterface interface {
	Push(data int)
	Pop()
}

type Queue struct {
	list *List
}

func NewQueue(cap int) *Queue {
	return &Queue{
		list: &List{},
	}
}

func (q *Queue) Push(data int) {
	node := &Node{
		Data: data,
	}
	// s.list.Append(node)
}

//TODO: write Pop method
//TODO: create main.go
