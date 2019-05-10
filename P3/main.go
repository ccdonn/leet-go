package main

import (
	"log"
	"strings"
)

func main() {
	s := "abcabcbb"
	max := 0
	var str strings.Builder

	for _, r := range s {
		v := str.String()
		if i := strings.Index(v, string(r)); i > -1 {
			if str.Len() > max {
				max = str.Len()
			}
			str.Reset()
			str.WriteString(v[i+1 : len(v)])
		}
		str.WriteRune(r)
	}
	if len(str.String()) > max {
		max = len(str.String())
	}
	// return max
	log.Println(max)
}
