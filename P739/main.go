package main

import "log"

func main() {
	log.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(T []int) []int {
	r := make([]int, len(T))
	c := 0

	for i, v := range T {
		d := 0
		for _, v2 := range T[i+1 : len(T)] {
			d++
			if v2 > v {
				r[c] = d
				break
			}
		}
		d = 0
		c++
	}

	return r
}
