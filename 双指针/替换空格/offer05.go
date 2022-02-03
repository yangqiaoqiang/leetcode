package main

import "fmt"

func main() {
	s := "hello world !"
	fmt.Println(replaceSpace(s))
}
func replaceSpace(s string) string {
	b := []byte(s)
	l := len(b) - 1
	flag := 0
	for _, v := range b {
		if v == ' ' {
			flag++
		}
	}
	flag = flag * 2
	b1 := make([]byte, flag)
	b = append(b, b1...)
	r := len(b) - 1
	for l >= 0 {
		if b[l] != ' ' {
			b[r] = b[l]
			r--
			l--
		} else {
			b[r] = '0'
			b[r-1] = '2'
			b[r-2] = '%'
			l--
			r = r - 3
		}
	}
	return string(b)
}
