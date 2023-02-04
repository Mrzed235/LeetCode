package main

//func main() {
//	var input string
//	fmt.Scanln(&input)
//	rs := buildToArr(input)
//	list := CreateList()
//	for _, r := range rs {
//		list.Append(r)
//	}
//	//sortList(list.Header)
//}

func sortList(head *ListNode) *ListNode {
	return nil
}

func sortA(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sortA(head, mid), sortA(mid, tail))
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val.(int) <= temp2.Val.(int) {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}
