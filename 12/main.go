package main

func main() {
	const max = 100000000
	//primesCount := make([]map[int]int, max)
	//
	//for i := 2; i < max; i++ {
	//	if primesCount[i] == nil {
	//		primesCount[i] = map[int]int { i: 1 }
	//		for j := 2 * i; j < max; j += i {
	//			if primesCount[j] == nil {
	//				primesCount[j] = map[int]int {}
	//			}
	//			primesCount[j][i] = primesCount[j/i][i] + 1
	//		}
	//	}
	//}

	for i := 1; ; i++ {
		div := 0
		for j := 1; j <= (i + 1)/2; j++ {
			if (i + 1)/2 % j == 0 {
				
			}
		}
	}
}
