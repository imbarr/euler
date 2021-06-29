package main

import (
	"time"
	"log"
)

func main() {
	const number = 600851475143
	const step = 500

	isComposite := make([]bool, step+1)
	left := number
	iteration := 1

	start := time.Now()

	for {
		limit := step*iteration

		for i := 2; i < limit; i++ {
			if !isComposite[i] {
				for j := 2 * i; j <= limit; j += i {
					isComposite[j] = true
				}
			}
		}

		for i := 2; i <= limit; i++ {
			if !isComposite[i] && left % i == 0 {
				for left % i == 0 {
					left = left / i
				}
				if left == 1 {
					log.Printf("Time: %s", time.Since(start))
					println(i)
					return
				}
			}
		}

		log.Printf("Iteration %d, left = %d", iteration, left)

		iteration++
		isComposite = make([]bool, iteration*step + 1)
	}
}
