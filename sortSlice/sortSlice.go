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
	arrCopy := append([]int(nil), arr...)
	sorted := make([]int, 0, len(arrCopy))

	for len(arrCopy) > 0 {
		sIndex := smallIndex(arrCopy)
		sorted = append(sorted, arrCopy[sIndex])
		arrCopy = append(arrCopy[:sIndex], arrCopy[sIndex+1:]...)
	}

	return sorted
}
