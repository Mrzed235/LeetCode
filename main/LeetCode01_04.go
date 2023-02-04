package main

/*
/给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
回文串不一定是字典当中的单词

输入："tactcoa"
输出：true（排列有"tacocat"、"atcocta"，等等）

输入："AaBb//a"
输出：false（区分大小写/也算字符）
*/

func canPermutePalindrome(s string) bool {
	//定义数组来保存字符,数组下标0存a出现的次数，依此类推
	rs := make([]int, 1000)
	bs := []byte(s)
	for _, str := range bs {
		rs[str]++
	}
	/*
			/rs中有效数字指出现次数不为0的字符位
			如果有效字符个数为1，则对字符出现的次数无要求
			有效字符个数为2，则需要1种字符的次数为奇数，一种为字符
			3->，则需要2种字符的出现的次数为奇数，1种为偶数，以此类推
		n->场景时候 n-1种字符的次数为odd，1种为奇数
	*/
	var flag = 0      //此变量记录有效字符的个数
	var odd, even int //odd记录有效字符的偶数次数的个数，even记录有效字母的奇数次数的个数
	for _, r := range rs {
		if r != 0 {
			flag++
			if r%2 == 0 {
				odd++
			} else {
				even++
			}
		}
	}
	if flag == 1 || odd == flag {
		return true
	}
	if odd == flag-1 && even == 1 {
		return true
	}
	return false
}

func canPermutePalindromeB(s string) bool {
	//定义数组来保存字符,数组下标0存a出现的次数，依此类推
	rs := make([]int, 1000)
	bs := []byte(s)
	for _, str := range bs {
		rs[str]++
	}
	/*
		有效字母的奇数次数的个数 <= 1符合条件
	*/
	var even int //even记录有效字母的奇数次数的个数
	for _, r := range rs {
		if r%2 == 1 {
			even++
		}
	}
	if even <= 1 {
		return true
	}
	return false
}

//func main() {
//	var str string
//	fmt.Scanln(&str)
//	fmt.Println()
//	fmt.Println(canPermutePalindrome(str))
//	fmt.Println(canPermutePalindromeB(str))
//}
