package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
LetCode2:给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。你可以假设除了数字 0 之外，这两个数都不会以 0开头。
示例 ：
2->4->3
5->6->4
输出：
7->0->8

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
链表的解题思路，创建链表对其进行遍历时要用两个标识分别记录头和移动标识位，遍历数据处理结束后就不会丢掉头的地址。

*/

func main() {
	var l1 string
	var l2 string
	_, _ = fmt.Scanln(&l1)
	_, _ = fmt.Scanln(&l2)
	arr1s := buildArray(l1)
	arr2s := buildArray(l2)
	header1 := CreateList()
	header2 := CreateList()
	for _, arr1 := range arr1s {
		header1.Append(arr1)
	}
	for _, arr2 := range arr2s {
		header2.Append(arr2)
	}
	head := addTwoNumbers(header1.Header, header2.Header)
	Skinhead(head)
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	//创建新链表头和标识位
	var tail *ListNode
	//定义进位符号 carry
	var carry = 0
	for l1 != nil || l2 != nil {
		//定义变量保存两个链表中的值
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.val.(int)
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.val.(int)
			l2 = l2.Next
		}
		//定义sum保存两个链表节点的和
		sum := n1 + n2 + carry
		carry = sum / 10 //如果之和满足进位，求出进的位数6+4 =10，进位1 余数0，
		sum = sum % 10   //如果之和满足10会进位，求出进位后的余数，后边要存入L3链表的节点中
		// 则保存0进1给下一个节点求和用,如果之和不大于等于10则carry就=0了，上轮节点计算符合进位的carry就不会被残留，就会被计入当前节点的sum值
		if head == nil {
			//新链表中无数据，所以将上述的计算和sum当值存入新的node，如果carry不=0则进位
			head = &ListNode{val: sum}
			tail = head //将head的地址赋值给tail进行遍历
		} else {
			tail.Next = &ListNode{val: sum}
			tail = tail.Next
		}
	}
	//如果循环完毕后carry!=0的话，则需要创建新节点将carry保存，并链接到之前的链表尾端（carry!=0的情况，两个链表最后一个节点之和大于10，符合进位条件）
	if carry > 0 {
		tail.Next = &ListNode{val: carry}
	}
	return
}

func Skinhead(head *ListNode) {
	current := head
	i := 1
	for current.Next != nil {
		fmt.Printf("第%d的节点是%d,下一个节点的地址是%p\n", i, current.val, current.Next)
		current = current.Next
		i++
	}
	fmt.Printf("第%d的节点是%d,下一个节点的地址是%p\n", i, current.val, current.Next)

}

func buildArray(str string) []int {
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	arr := strings.Split(str, ",")
	rs := make([]int, 0, len(arr))
	for _, s := range arr {
		r, _ := strconv.Atoi(s)
		rs = append(rs, r)
	}
	return rs
}
