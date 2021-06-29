package main

func factorial(x int) int {
	r := 1
	for i := 1; i <= x; i++ {
		r *= i
	}
	return r
}

func C(n, r int) int {
	prod := 1
	for i := n - r + 1; i <= n; i++ {
		prod *= i
	}

	return prod / factorial(r)
}

func main() {
	res := 0

	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			if C(n, r) > 1000000 {
				res += n - 2*r + 1
				break
			}
		}
	}

	println(res)
}
