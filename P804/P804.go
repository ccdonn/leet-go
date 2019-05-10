package main

var arr = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	hmap := make(map[string]int)
	for _, word := range words {
		s := ""
		for i := 0; i < len(word); i++ {
			s += arr[int(word[i])-97]
		}
		hmap[s] = 1
	}
	return len(hmap)
}
