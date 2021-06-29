package main

func primeMap(max int) map[int][]int {
	isComp := make([]bool, max + 1)
	res := make(map[int][]int)

	for i := 2; i <= max; i++ {
		if !isComp[i] {
			res[i] = make([]int, max / i - 1)
			ptr := 0
			for j := 2*i; j <= max; j+=i {
				isComp[j] = true
				res[i][ptr] = j
				ptr++
			}
		}
	}

	return res
}

func main() {
	const max = 1000000

	pMap := primeMap(max)
	maxRes := 0.0

	for n := 2; n <= max; n++ {
		primes := pMap[n]
		for _, prime := range primes {

		}
	}
}
