package main

import "fmt"

func main() {
	s := "string"
	k := 2
	fmt.Println(reverseStr(s, k))
}
func reverseStr(s string, k int) string {
	ss := []byte(s)
	for i := 0; i <= len(ss); {
		l := i
		r := i + k - 1
		if r > len(ss)-1 {
			r = len(ss) - 1
		}
		for l < r {
			ss[l], ss[r] = ss[r], ss[l]
			l++
			r--
		}
		i = i + 2*k
	}
	s = string(ss)
	return s
}
