package main

import (
	"time"
	"log"
)

func main() {
	fib := make([]int, 0)
	prev, next := 1, 2

	start := time.Now()

	for prev <= 4000000 {
		fib = append(fib, prev)
		newNext := prev + next
		prev, next = next, newNext
	}

	log.Printf("Array fill: %s", time.Since(start))
	arrayFill := time.Now()

	sum := 0
	for _, el := range fib {
		if el % 2 == 0 {
			sum += el
		}
	}

	log.Printf("Sum calc: %s", time.Since(arrayFill))
	println(sum)
}