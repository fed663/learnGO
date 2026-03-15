package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x < 10 {
		return true
	}

	arr := []byte(strconv.Itoa(x))
	head, tail := 0, len(arr)-1
	for head < tail {
		if arr[head] != arr[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}
