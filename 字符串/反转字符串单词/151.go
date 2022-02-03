package main

import "fmt"

func main() {
	s := " hello world "
	fmt.Println(reverseWords(s))
}
func reverseWords(s string) string {
	b := []byte(s)
	slow, fast := 0, 0
	if len(b) <= 1 {
		return string(b)
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
		fast++
		slow++
	}
	if slow > 0 && b[slow-1] == ' ' {
		b = b[:slow-1]
	} else {
		b = b[:slow]
	}
	reverse(&b, 0, len(b)-1)
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}
func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}
