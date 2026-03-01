package main

import (
	"errors"
)

var ErrEmptySlice = errors.New("срез пустой")

func max(arr []int) (int, error) {

	if len(arr) == 0 {
		return 0, ErrEmptySlice
	}
	if len(arr) == 1 {
		return arr[0], nil
	}

	elMax, err := max(arr[1:])
	if err != nil {
		return 0, err
	}
	if arr[0] > elMax {
		return arr[0], nil
	}
	return elMax, nil
}
