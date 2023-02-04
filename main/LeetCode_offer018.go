package main

import (
	"strings"
)

//给定一个字符串s，验证字符串s是否是回文字符串，可以忽略大小写
/*
/输入: s = "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串

输入: s = "race a car"
输出: false
解释："raceacar" 不是回文串
*/

//func main() {
//	var str string
//	fmt.Scanln(&str)
//	fmt.Println(isPalindrome(str))
//}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	strs := []byte(s)
	chars := make([]byte, 0, 0)
	for _, s := range strs {
		if (s >= byte('0') && s <= byte('9')) || (s >= byte('a') && s <= byte('z')) {
			chars = append(chars, s)
		}
	}
	n := len(chars)
	for i, v := range chars[:n/2] {
		if v != chars[n-1-i] {
			return false
		}
	}
	return true
}
