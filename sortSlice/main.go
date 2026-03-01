package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	testSlice := []int{228, 23, 664, 7, 0, 52, 8, 8, 5, 1337}
	sortedSlice := sortSlice(testSlice)
	_, _ = fmt.Fprintln(os.Stdout, sortedSlice)

	testSlice2 := []int{2, 4, 6, 8}
	sumTest := sum(testSlice2)
	_, _ = fmt.Fprintln(os.Stdout, sumTest)

	countTest := count(testSlice2)
	_, _ = fmt.Fprintln(os.Stdout, countTest)

	maxTest, err := max([]int{})
	if err != nil {
		if errors.Is(err, ErrEmptySlice) {
			fmt.Println("expected error:", err)
		} else {
			fmt.Println("unexpected error:", err)
			return
		}
	} else {
		_, _ = fmt.Fprintln(os.Stdout, maxTest)
	}

	fastSortTest := fastSortSlice(testSlice)
	_, _ = fmt.Fprintln(os.Stdout, fastSortTest)
}
