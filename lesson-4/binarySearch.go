package main

import (
	"errors"
	"fmt"
)

//TODO: попробовать реализовать через интерфейс и методы для любой сущности, которая имеет длину и отсортирована не привязывясь только к массиву

func binarySearch(arr []int32, reqNum int32) (reqIdx int32, err error) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == reqNum {
			reqIdx = int32(mid)
			return reqIdx, nil
		}

		if arr[mid] > reqNum {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	err = errors.New("Number not found")
	return reqNum, err
}

func main() {
	arr := []int32{1, 2, 3, 4, 5}

	fmt.Println(binarySearch(arr, 2)) //TODO: добавить ввод с клавиатуры или через аргументы командной строки
}
