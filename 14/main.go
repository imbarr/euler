package main

func chainLen(start int) int {
	count := 1
	for cur := start; cur > 1; {
		count++
		if cur % 2 == 0 {
			cur /= 2
		} else {
			cur = 3*cur + 1
		}
	}

	return count
}

func main() {
	max := 0
	maxI := -1

	for i := 1000000; i > 0; i-- {
		l := chainLen(i)
		if l > max {
			max = l
			maxI = i
		}
	}

	println(maxI)
}
