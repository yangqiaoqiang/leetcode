package main

import "fmt"

func main() {
	s := "abcdefg"
	n := 2
	fmt.Println(reverseLeftWords(s, n))
}
func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	reverse(b, 0, len(b)-1)
	reverse(b, 0, len(b)-1-n)
	reverse(b, len(b)-n, len(b)-1)
	return string(b)
}
func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
