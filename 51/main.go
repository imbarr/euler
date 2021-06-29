package main

import "math"

func factorial(x int) int {
	r := 1
	for i := 1; i <= x; i++ {
		r *= i
	}
	return r
}

func getCombinations(arr []int) [][]int {
	res := make([][]int, int(math.Pow(2, float64(len(arr)))) - 1)
	pointer := 0

	for size := 1; size <= len(arr); size++ {
		sub := getSizeCombinations(arr, size)
		copy(res[pointer:], sub)
		pointer += len(sub)
	}

	return res
}

func getSizeCombinations(arr []int, size int) [][]int {
	if size == 0 {
		return [][]int { make([]int, 0) }
	}

	length := len(arr)
	posSize := factorial(length) / factorial(length - size) / factorial(size)
	positions := make([][]int, posSize)
	pointer := 0

	for i := 0; i <= length - size; i++ {
		sub := getSizeCombinations(arr[i+1:], size - 1)
		for _, el := range sub {
			newPos := make([]int, size)
			newPos[0] = arr[i]
			copy(newPos[1:], el)
			positions[pointer] = newPos
			pointer++
		}
	}

	return positions
}

func compositeNumbers(max int) []bool {
	isComposite := make([]bool, max + 1)

	for i := 2; i <= max; i++ {
		if !isComposite[i] {
			for j := 2 * i; j <= max; j += i {
				isComposite[j] = true
			}
		}
	}

	return isComposite
}

func toDigitArray(num int) []byte {
	digitCount := int(math.Log10(float64(num))) + 1
	digits := make([]byte, digitCount)
	for i, _ := range digits {
		digits[i] = byte(num % int(math.Pow10(digitCount - i)) / int(math.Pow10(digitCount - i - 1)))
	}
	return digits
}

func fromDigitArray(arr []byte) int {
	sum := 0
	for i, digit := range arr {
		sum += int(digit) * int(math.Pow10(len(arr) - i - 1))
	}
	return sum
}


func main() {
	const max = 6
	const special = 1
	const count = 1

	composites := compositeNumbers(int(math.Pow10(max)))

	for num := 2; num < len(composites); num++ {
		if composites[num] {
			continue
		}

		digits := toDigitArray(num)
		specialDigits := make([]int, 0)
		for i, _ := range digits {
			if digits[i] == special {
				specialDigits = append(specialDigits, i)
			}
		}

		if len(specialDigits) == 0 {
			continue
		}

		combinations := getCombinations(specialDigits)

		for _, comb := range combinations {
			primeCount := 0
			for repl := 1; repl <= 9; repl++ {
				for _, specialDigit := range comb {
					digits[specialDigit] = byte(repl)
				}
				if !composites[fromDigitArray(digits)] {
					primeCount++
				}
			}

			if primeCount >= 8 {
				println(num)
				for _, el := range comb {
					print(el)
					print(" ")
				}
				println()
				return
			}
		}
	}

	println("Not found")
}
