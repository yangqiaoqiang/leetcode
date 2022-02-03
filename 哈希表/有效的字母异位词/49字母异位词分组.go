package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"aa", "ab", "ba"}
	fmt.Println(groupAnagrams(a))
}
func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := make(map[string][]string, len(strs))
	for _, v := range strs {
		s := []byte(v)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) // 排序
		key := string(s)                                          // 异位词排序后一定是相同的
		// map的键是排序后的字符串,值是字符串切片,用于存储分组
		m[key] = append(m[key], v)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
