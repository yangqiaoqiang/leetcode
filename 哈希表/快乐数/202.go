package main

import "fmt"

func main() {
	i := 19
	fmt.Println(isHappy(i))
}
func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		m[n] = true
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		n = sum
	}
	return n == 1
}
