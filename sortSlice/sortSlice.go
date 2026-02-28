package main

func smallIndex(arr []int) int {
	smallest := arr[0]
	indexSmallest := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			indexSmallest = i
		}
	}

	return indexSmallest
}

func sortSlice(arr []int) []int {
	sorted := make([]int, 0, len(arr))

	for len(arr) > 0 {
		sIndex := smallIndex(arr)
		sorted = append(sorted, arr[sIndex])
		arr = append(arr[:sIndex], arr[sIndex+1:]...)
	}

	return sorted
}
