package models

type StackInterface interface {
	Push(data int)
	Pop()
}

type Stack struct {
	list *List
}

func NewStack(cap int) *Stack {
	return &Stack{
		list: &List{},
	}
}

func (s *Stack) Push(data int) {
	node := &Node{
		Data: data,
	}
	s.list.Append(node)
}

func (s *Stack) Pop() int {
	if s.list.Len() == 0 {
		return 0
	}

	//TODO: why nil pointer dereference?
	elem := s.list.Tail().Data
	s.list.Delete(s.list.Tail())
	return elem
}
