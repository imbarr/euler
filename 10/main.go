package main

func main() {
	const max = 2000000
	isComposite := make([]bool, max)

	for i := 2; i < max; i++ {
		if !isComposite[i] {
			for j := 2 * i; j < max; j += i {
				isComposite[j] = true
			}
		}
	}

	sum := 0
	for i := 2; i < max; i++ {
		if !isComposite[i] {
			sum += i
		}
	}

	println(sum)
}
