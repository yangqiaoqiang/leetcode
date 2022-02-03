// 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，
// 请你判断二者是否相等。# 代表退格字符。

// 如果相等，返回 true ；否则，返回 false 。
// 注意：如果对空文本输入退格字符，文本继续为空。
package main

import "fmt"

func main() {
	s := "abcd"
	t := "abcd"
	fmt.Println(backspaceCompare(s, t))
}
func backspaceCompare(s, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
