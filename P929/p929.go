package main

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)
	for _, v := range emails {
		// fmt.Println(v)
		if m[v] == false {
			// fmt.Println(m[v])
			// fmt.Println(strings.Index(v, "@"))
			// fmt.Println(v[0:strings.Index(v, "@")])
			// fmt.Println(strings.Index(v[0:strings.Index(v, "@")], "+"))
			// fmt.Println(v[0:strings.Index(v, "+")])
			// fmt.Println(strings.Replace(v[0:strings.Index(v, "+")], ".", "", -1))
			// fmt.Printf("%s%s\n", strings.Replace(v[0:strings.Index(v, "+")], ".", "", -1), v[strings.Index(v, "@"):len(v)])
			if strings.Index(v, "+") > 0 {
				m[strings.Replace(v[0:strings.Index(v, "+")], ".", "", -1)+v[strings.Index(v, "@"):len(v)]] = true
			} else {
				m[strings.Replace(v, ".", "", -1)+v[strings.Index(v, "@"):len(v)]] = true
			}

			// m[v] = true
		}

	}
	return len(m)
}
