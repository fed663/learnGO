package main

import (
	"fmt"
)

func fastSortSlice(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return []int{arr[1], arr[0]}
		} else {
			return []int{arr[0], arr[1]}
		}
	} else {
		elem := len(arr) / 2
		lowArr := make([]int, 0, len(arr))
		higArr := make([]int, 0, len(arr))

		for i, v := range arr {
			if i == elem {
				continue
			}
			if v > arr[elem] {
				higArr = append(higArr, v)
			} else {
				lowArr = append(lowArr, v)
			}
		}

		fmt.Println(lowArr)
		fmt.Println(higArr)
	}
	return nil
}
