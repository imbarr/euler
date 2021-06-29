package main

func main() {
	const max = 1000000
	isComposite := make([]bool, max)
	count := 0

	for i := 2; i < max; i++ {
		if !isComposite[i] {
			count++
			if count == 10001 {
				println(i)
				return
			}
			for j := 2 * i; j < max; j += i {
				isComposite[j] = true
			}
		}
	}
}
