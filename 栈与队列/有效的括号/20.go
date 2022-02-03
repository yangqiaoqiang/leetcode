package main

func main() {

}
func isValid(s string) bool {
	hash := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	flag := []byte{}
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			flag = append(flag, s[i])
		} else if len(flag) > 0 && flag[len(flag)-1] == hash[s[i]] {
			flag = flag[:len(flag)-1]
		} else {
			return false
		}
	}
	if len(flag) > 0 {
		return false
	}
	return true
}
