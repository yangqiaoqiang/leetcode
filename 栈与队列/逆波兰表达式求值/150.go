package main

import "strconv"

func main() {

}
func evalRPN(tokens []string) int {
	s := []int{}
	for i := 0; i <= len(tokens)-1; i++ {
		flag, err := strconv.Atoi(tokens[i])
		if err == nil {
			s = append(s, flag)
		} else {
			n2, n1 := s[len(s)-1], s[len(s)-2]
			s = s[:len(s)-2]
			switch tokens[i] {
			case "+":
				s = append(s, n1+n2)
			case "-":
				s = append(s, n1-n2)
			case "*":
				s = append(s, n1*n2)
			case "/":
				s = append(s, n1/n2)
			}
		}
	}
	return s[0]
}
