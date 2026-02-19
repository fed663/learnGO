package main

import (
	"fmt"
	"os"
)

func main() {
	nums := make([]int, 128)

	for i := 1; i <= 128; i++ {
		nums[i-1] = i
	}

	var userNum int
	_, err := fmt.Fscan(os.Stdin, &userNum)
	if err != nil {
		panic(err)
	}

	fmt.Println(binarySearch(nums, userNum))
}

func binarySearch(arr []int, val int) int {
	low := 0
	high := len(arr) - 1
	steps := 0

	for low <= high {
		steps++
		mid := (low + high) / 2
		if arr[mid] == val {
			fmt.Printf("Пройдено шагов: %v\n", steps)
			return mid
		} else if arr[mid] < val {
			low = mid + 1
		} else if arr[mid] > val {
			high = mid - 1
		}
	}
	return -1
}
