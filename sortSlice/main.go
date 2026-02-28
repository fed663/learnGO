package main

import (
	"fmt"
	"os"
)

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

func main() {
	testSlice := []int{228, 23, 664, 7, 0, 52, 8, 8, 5, 1337}
	sortedSlice := sortSlice(testSlice)
	_, err := fmt.Fprintln(os.Stdout, sortedSlice)
	if err != nil {
		panic(err)
	}

	testSlice2 := []int{2, 4, 6, 8}
	sumTest := sum(testSlice2)
	_, err = fmt.Fprintln(os.Stdout, sumTest)
	if err != nil {
		panic(err)
	}

	countTest := count(testSlice2)
	_, err = fmt.Fprintln(os.Stdout, countTest)
	if err != nil {
		panic(err)
	}

	maxTest := max(testSlice2)
	_, err = fmt.Fprintln(os.Stdout, maxTest)
	if err != nil {
		panic(err)
	}
}
