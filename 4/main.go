package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	r := make([]rune, len(s))
	for i, c := range s {
		r[len(s) - i - 1] = c
	}

	return string(r) == s
}

func main() {
	max := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			p := i*j
			if isPalindrome(p) && p > max {
				max = p
			}
		}
	}

	println(max)
}
