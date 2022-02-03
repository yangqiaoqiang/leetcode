package main

import "fmt"

func main() {
	s := "string"
	t := "strngg"
	fmt.Println(isAnagram(s, t))
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := make(map[byte]int)
	for i := range s {
		hash[s[i]]++
		hash[t[i]]--
	}
	for _, v := range hash {
		if v > 0 {
			return false
		}
	}
	return true
}
