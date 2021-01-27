package main

import "errors"

type data []int32

func binarySearch(arr data, reqNum int32) (reqIdx int32, err error) {
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
