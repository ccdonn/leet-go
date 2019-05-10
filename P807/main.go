package main

import (
	"log"
)

var sample = [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}

func main() {
	log.Println(maxIncreaseKeepingSkyline(sample))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	length := len(grid)
	h := make([]int, length)
	v := make([]int, length)
	ht := make([]int, length)
	vt := make([]int, length)
	c := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			ht[c] = grid[i][j]
			vt[c] = grid[j][i]
			c++
		}
		h[i] = findMax(ht...)
		v[i] = findMax(vt...)
		c = 0
	}

	log.Println("h", h)
	log.Println("v", v)

	r := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if grid[i][j] < h[i] && grid[i][j] < v[j] {
				r += findMin(h[i], v[j]) - grid[i][j]
			}
		}
	}

	return r
}

func findMax(args ...int) int {
	// sort.Ints(args)
	// return args[len(args)-1]
	max := -1
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}

func findMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
