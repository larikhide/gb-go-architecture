package main

import (
	"fmt"

	"./models"
)

func main() {
	queue := models.NewQueue(2)
	queue.Push(5)
	queue.Push(6)
	queue.Push(7)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

	stack := models.NewStack(2)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
