package main

import (
	"fmt"
	"strings"
)

/*
/无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
  请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串

*/
func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	var ans int                 //定义ans保存当前窗口中的最长无重复字符串长度
	cwindow := make([]int, 128) //定义数组来保存右边窗口所在字符的下标
	for i := range cwindow {
		cwindow[i] = -1
	}
	s = strings.Trim(s, "\"")
	bs := []byte(s)
	left, right := 0, 0
	for right = 0; right < len(bs); right++ {
		if cwindow[bs[right]] != -1 {
			left = max(left, cwindow[bs[right]]+1) //加max是因为判断当前字符出现的下标在left左边，则left选当前值，因为left不能左移
		}
		ans = max(ans, right-left+1) //求当前窗口不重复字符串的最大值且保存到ans进行下轮比较
		cwindow[bs[right]] = right
	}
	return ans
}

func max(l, r int) int {
	if l < r {
		return r
	}
	return l
}
