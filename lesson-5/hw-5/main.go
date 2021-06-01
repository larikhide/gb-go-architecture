package main

import (
	"fmt"

	"./models"
)

func main() {

	queue := models.NewQueue(2)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

	stack := models.NewStack(2)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Push(7)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

}
