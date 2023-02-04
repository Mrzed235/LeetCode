package main

import "fmt"

func update(a []int) {
	a[1] = 10
	fmt.Printf("%p\n", &a)
}

//func main() {
//	a := []int{0, 1, 2, 3, 4}
//	fmt.Println(a)
//	fmt.Printf("%p\n", &a)
//	update(a)
//	fmt.Printf("%p\n", &a)
//	fmt.Println(a)
//}
