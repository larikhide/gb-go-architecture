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
	// s.list.Append(node)
}

//TODO: write Pop method
