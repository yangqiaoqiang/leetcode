package main

import "fmt"

func main() {
	s := "hello world !"
	fmt.Println(replaceSpace(s))
}
func replaceSpace(s string) string {
	ss := []byte(s)

	res := []byte{}
	for i := 0; i <= len(ss)-1; i++ {
		if ss[i] != ' ' {
			res = append(res, ss[i])
		} else {
			res = append(res, []byte("%20")...)
		}
	}
	ress := string(res)
	return ress
}
