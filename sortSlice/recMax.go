package main

func max(arr []int) int {
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		} else {
			return arr[1]
		}
	}
	elMax := max(arr[1:])
	if arr[0] > elMax {
		return arr[0]
	} else {
		return elMax
	}
}
