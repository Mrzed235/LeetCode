package main

import "fmt"

/*
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false
输入：head = [1,2,2,1]  1->2->2->1
输出：true
输入：head = [1,2]
输出：false
*/

func isPalindromeA(head *ListNode) bool {
	vals := []int{}
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val.(int))
	}
	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}
func isPalindromeB(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	var prev *ListNode = nil
	for fast != nil && fast.Next != nil {
		prev = slow //当fast遍历到结尾时，slow会指向中点，此时的prev保存了slow移动前的位置
		fast = fast.Next.Next
		slow = slow.Next
	}
	prev.Next = nil //prev的作用就在与记录位置来断开链表，此时slow为头的链表保存了后半段链表的head值
	//反转后半段链表
	var head2 *ListNode = nil
	for slow != nil {
		tmp := slow.Next  //定义tmp记录每一轮的slow的后边的节点的头部
		slow.Next = head2 //第一轮的原链表的的一个元素指向nil,
		head2 = slow      //
		slow = tmp
	}
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

func main() {
	var input string
	fmt.Scanln(&input)
	rs := buildToArr(input)
	list := CreateList()
	for _, r := range rs {
		list.Append(r)
	}
	fmt.Println(isPalindromeA(list.Header))
	fmt.Println(isPalindromeB(list.Header))
}
