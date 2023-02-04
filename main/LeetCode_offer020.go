package main

/*
/给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"

输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
*/

//todo：待做
func buildSubstring(s string) []string {
	//划分子串的粒度为1--->len(s)-1
	ans := make([]string, 0)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			ans = append(ans, s[i:j])
		}
	}
	return ans
}

func countSubstrings(s string) int {
	rs := buildSubstring(s)
	var count = 0
	for _, r := range rs {
		if isPalindrome(r) {
			count++
		}
	}
	return count
}

func countSubstringsB(s string) (ans int) {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		ans++
	} // len=1的字符，即s[i]自成回文子串
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			ans++
		}
	} // len=2的字符，若s[i]==s[i+1]则自成回文子串
	for LEN := 3; LEN <= n; LEN++ { // len=3、4、5...n的字符，若s[i]==s[i+LEN] && dp[i][i+LEN] == true 则自成回文子串
		for i := 0; i <= n-LEN; i++ { // i in range [0, n-LEN]
			j := i + LEN - 1 // j in range [LEN, n-1]
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
				if dp[i][j] {
					ans++
				}
			}
		}
	}
	return
}

//func main() {
//	var s string
//	fmt.Scanln(&s)
//
//	//fmt.Println(countSubstrings(s))
//	fmt.Println(countSubstringsB(s))
//
//}
