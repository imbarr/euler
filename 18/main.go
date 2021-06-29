package main

import (
	"math"
	"strconv"
	"strings"
)

func dijkstra(graph [][]int, src, dst int) int {
	n := len(graph)
	dist := make([]int, n)
	visited := make([]bool, n)

	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[src] = 0

	for {
		current := -1
		minDist := math.MaxInt64
		for i, d := range dist {
			if d < minDist && !visited[i] {
				current = i
				minDist = d
			}
		}

		if current == -1 {
			return dist[dst]
		}

		for vert, d := range graph[current] {
			if d != -1 {
				newDist := dist[current] + d
				if newDist < dist[vert] {
					dist[vert] = newDist
				}
			}
		}

		visited[current] = true
	}
}

const raw = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

func main() {
	lines := strings.Split(raw, "\n")
	size := (len(lines) + 1) * len(lines) / 2
	graph := make([][]int, size + 2)
	for i := range graph {
		graph[i] = make([]int, size + 2)
		for j := range graph[i] {
			graph[i][j] = -1
		}
	}
	first, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}
	graph[0][1] = 100 - first

	current := 1
	for _, line := range lines[1:] {
		chars := strings.Split(line, " ")
		for j, char := range chars {
			current++
			length, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			if j != 0 {
				graph[current - len(chars)][current] = 100 - length
			}
			if j != len(chars) - 1 {
				graph[current - len(chars) + 1][current] = 100 - length
			}
		}
	}

	last := len(graph) - 1
	for i := last - len(lines); i < last; i++ {
		graph[i][last] = 0
	}

	//for _, l := range graph {
	//	for _, el := range l {
	//		print(el, "\t")
	//	}
	//	println()
	//}

	res := dijkstra(graph, 0, last)
	println(100*len(lines) - res)
}
