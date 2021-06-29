package main

func getValues(f func(int) int) []int {
	res := make([]int, 0)
	for i := 0; ; i++ {
		value := f(i)
		if value >= 10000 {
			return res
		}
		if value >= 1000 {
			res = append(res, value)
		}
	}
}

func generateValues(channel chan []int, from, to, count int) {
	if count == 0 {
		channel <- make([]int, 0)
		close(channel)
		return
	}

	for i := from; i < to; i++ {
		sub := make(chan []int)
		go generateValues(sub, from, to, count-1)
		for arr := range sub {
			newArr := make([]int, count)
			newArr[0] = i
			copy(newArr[1:], arr)
			channel <- newArr
		}
	}

	close(channel)
}

func join(x, y int) int {
	return x * 100 + y
}

func contains(slice []int, el int) bool {
	for _, x := range slice {
		if x == el {
			return true
		}
	}
	return false
}

func main() {
	sets := [][]int {
		getValues(func(n int) int {
			return n * (n + 1) / 2
		}),
		getValues(func(n int) int {
			return n * n / 2
		}),
		getValues(func(n int) int {
			return n * (3*n - 1) / 2
		}),
		getValues(func(n int) int {
			return n * (2*n - 1)
		}),
		getValues(func(n int) int {
			return n * (5*n - 3) / 2
		}),
		getValues(func(n int) int {
			return n * (3*n - 2)
		}),
	}

	c := make(chan []int)
	go generateValues(c, 10, 100, 6)

	//for arr := range c {
	//	_0 := join(arr[0], arr[1])
	//	_1 := join(arr[1], arr[2])
	//	_2 := join(arr[2], arr[3])
	//	_3 := join(arr[3], arr[4])
	//	_4 := join(arr[4], arr[5])
	//	_5 := join(arr[5], arr[0])
	//
	//	if contains(sets[0], _0) &&
	//		contains(sets[0], _1) &&
	//		contains(sets[0], _2) &&
	//		contains(sets[0], _3) &&
	//		contains(sets[0], _4) &&
	//		contains(sets[0], _5) {
	//
	//		println(_0 + _1 + _2 + _2 + _4 + _5)
	//		return
	//	}
	//}

	println(len(sets[0]))
	println(len(sets[1]))
	println(len(sets[2]))
	println(len(sets[3]))
	println(len(sets[4]))
	println(len(sets[5]))

	println("Not found")
}
