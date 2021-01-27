package main

type data []int32

func binarySearch(arr data, reqNum int32) (reqIdx int, err error) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == reqNum {
			reqIdx = mid
			return reqIdx, nil
		}

		if arr[mid] > reqNum {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	// как вернуть ошибку в строке "not found"?
	return nil, err
}
