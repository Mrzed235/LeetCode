package main

import "fmt"

type ListNode struct {
	val  interface{}
	Next *ListNode
}

type LList struct {
	Header *ListNode //指向第一个节点
	Length int
}

//func main() {
//	header := CreateList()
//	header.Append(1)
//	header.Append(2)
//	header.Append(3)
//	fmt.Printf("linknode lenth is %v\n", header.Length)
//	header.Scan()
//}

func CreateNode(v interface{}) *ListNode {
	return &ListNode{v, nil}
}
func CreateList() *LList {
	header := CreateNode(nil)
	return &LList{header, 0}
}

//往链表头增加一个节点,
func (l *LList) Add(data interface{}) {
	newNode := CreateNode(data)
	defer func() {
		l.Length++
	}()

	if l.Length == 0 {
		l.Header = newNode
	} else {
		newNode.Next = l.Header
		l.Header = newNode //头指针指向新加的
	}

}

//往链表尾加一个节点
func (l *LList) Append(data interface{}) {
	newNode := CreateNode(data)
	defer func() {
		l.Length++
	}()

	if l.Length == 0 {
		l.Header = newNode
	}
	if l.Length > 0 {
		current := l.Header
		for current.Next != nil { //循环找到最后一个节点
			current = current.Next
		}
		current.Next = newNode //把新节点地址给最后一个节点的Next
	}
}

//往i插入一个节点，后插
func (l *LList) Insert(i int, data interface{}) {
	defer func() {
		l.Length++
	}()

	if i >= l.Length {
		l.Append(data)
		return
	}
	newNode := CreateNode(data)
	//找到第i个节点pre和i+1个after节点
	j := 1
	pre := l.Header
	for j != i {
		pre = pre.Next
		j++
	}
	after := pre.Next //获取到i+1个节点
	//修改i节点，新节点的指针
	pre.Next = newNode
	newNode.Next = after
}

//删除第i个节点
func (l *LList) Delete(i int) {
	defer func() {
		l.Length--
	}()
	if i == 1 { //删除第一个节点，把header指向第二个节点即可
		l.Header = l.Header.Next
		return
	}
	//找到第i-1个节点，找到第i+1个节点，修改i-1的节点的next即可
	j := 0
	current := l.Header
	for j == i-1 {
		current = current.Next
		j++
	}
	after := current.Next.Next
	current.Next = after
}

//遍历链表，显示出来
func (l *LList) Scan() {
	current := l.Header
	i := 1
	for current.Next != nil {
		fmt.Printf("第%d的节点是%d,下一个节点的地址是%p\n", i, current.val, current.Next)
		current = current.Next
		i++
	}
	fmt.Printf("第%d的节点是%d,下一个节点的地址是%p\n", i, current.val, current.Next)

}
