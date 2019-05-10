package main

import (
	"strings"
)

func numJewelsInStones(J string, S string) int {
	sum := 0
	jewel := strings.Split(J, "")
	for _, v := range jewel {
		// fmt.Println(v)
		sum += strings.Count(S, v)
	}
	return sum
}
