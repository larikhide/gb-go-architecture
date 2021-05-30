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
	q.list.Append(node)
}

func (q *Queue) Pop() int {
	if q.list.Len() == 0 {
		return 0
	}

	elem := q.list.Head().Data
	q.list.Delete(q.list.Head())
	return elem
}

//TODO: create main.go
