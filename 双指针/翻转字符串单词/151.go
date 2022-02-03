package main

import "fmt"

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}
func reverseWords(s string) string {
	b := []byte(s)
	slow, fast := 0, 0
	if len(b) <= 1 {
		return s
	}
	for fast <= len(b)-1 {
		if b[fast] == ' ' {
			fast++
		} else {
			break
		}
	}
	for fast <= len(b)-1 {
		if fast > 0 && b[fast-1] == b[fast] && b[fast] == ' ' {
			fast++
			continue
		}
		b[slow] = b[fast]
		slow++
		fast++
	}
	if slow > 0 && b[slow-1] == ' ' {
		b = b[:slow-1]
	} else {
		b = b[:slow]
	}
	reverse(b, 0, len(b)-1)
	for i := 0; i <= len(b)-1; i++ {
		j := i
		for j <= len(b)-1 && b[j] != ' ' {
			j++

		}
		reverse(b, i, j-1)
		i = j
	}
	return string(b)
}
func reverse(b1 []byte, l, r int) {
	for l < r {
		b1[l], b1[r] = b1[r], b1[l]
		l++
		r--
	}
}
