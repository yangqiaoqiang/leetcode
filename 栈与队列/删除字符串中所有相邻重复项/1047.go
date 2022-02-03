package main

func main() {

}
func removeDuplicates(s string) string {
	flag := []byte{}
	for i := 0; i <= len(s)-1; i++ {
		if len(flag) > 0 && s[i] == flag[len(flag)-1] {
			flag = flag[:len(flag)-1]
		} else {
			flag = append(flag, s[i])
		}
	}
	return string(flag)
}
