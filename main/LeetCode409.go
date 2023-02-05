package main

/*
/409. 最长回文串  简单
给定一个包含大写字母和小写字母的字符串s，返回通过这些字母构造成的 最长的回文串。
在构造过程中，请注意 区分大小写 。比如"Aa"不能当做一个回文字符串。

输入:s = "abccccdd"
输出:7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

输入:s = "aaaaaccc"
输出:7
*/

//func main() {
//	var s string
//	fmt.Scanln(&s)
//	s = strings.Trim(s, "\"")
//	longestPalindrome(s)
//}

func longestPalindrome(s string) int {
	rs := make([]int, 128)
	byts := []byte(s)
	var ans = 0
	for _, byt := range byts {
		rs[byt]++
		if rs[byt] == 2 {
			ans += 2
			rs[byt] = 0
		}
	}
	if ans < len(s) {
		ans++
	}
	return ans
}
