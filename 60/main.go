package main

import "strconv"

func getSizeCombinations(arr []int, size int, c chan []int) {
	if size == 0 {
		c <- []int{}
		close(c)
		return
	}

	length := len(arr)

	for i := 0; i <= length - size; i++ {
		subChannel := make(chan []int)
		go getSizeCombinations(arr[i+1:], size - 1, subChannel)
		for sub := range subChannel {
			res := make([]int, size)
			res[0] = arr[i]
			copy(res[1:], sub)
			c <- res
		}
	}

	close(c)
}

func isCompositeArray(max int) ([]bool, int) {
	isComposite := make([]bool, max)
	count := 0

	for i := 2; i < max; i++ {
		if !isComposite[i] {
			count++
			for j := 2 * i; j < max; j += i {
				isComposite[j] = true
			}
		}
	}

	return isComposite, count
}



func main() {
	const max = 2000000
	const setCount = 5

	isComposite, _ := isCompositeArray(max)

	adjCount := make([]byte, max)
	adj := make([][]bool, max)

	for n := range isComposite {
		if !isComposite[n] {
			str := strconv.Itoa(n)

			for shift := 1; shift < len(str); shift++ {
				first, err := strconv.Atoi(str[:shift])
				if err != nil {
					panic(err)
				}
				second, err := strconv.Atoi(str[shift:])
				if err != nil {
					panic(err)
				}

				if !isComposite[first] && !isComposite[second] {
					if adj[first] == nil {
						adj[first] = make([]bool, max)
					}
					adj[first][second] = true
					if adj[second] != nil && adj[second][first] {
						adjCount[second]++
						adjCount[first]++
					}
				}
			}
		}
	}

	subgraph := make([]int, 1000)[:0]
	subgraphSet := make(map[int]bool, 1000)
	for i, v := range adjCount {
		if v >= setCount - 1 {
			subgraph = append(subgraph, i)
			subgraphSet[i] = true
		}
	}

	println(len(subgraph))

	t := make([]int, 446)
	c := make(chan []int)
	count := 0
	go getSizeCombinations(t, 5, c)
	for range c {
		count++
		if count % 1000000 == 0 {
			println(count)
		}
	}
	println(count)
}
