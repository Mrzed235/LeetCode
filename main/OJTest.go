/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
 * Description: 上机编程认证
 * Note: 缺省代码仅供参考，可自行决定使用、修改或删除
 * 只能import Go标准库
 */

/*
/输入：7
shenzhen ebg 5
shenzhen dbg 7
shenzhen cbg 5
shenzhen abg 6
shenzhen zbg 5
shenzhen bbg 4
beijing cbg 1
输出：
beijing cbg
shenzhen abg
shenzhen cbg
shenzhen dbg
shenzhen ebg
shenzhen zbg
输入：6
wuhan appdept 100
wuhan app 100
wuhan cpdept 100
wuhan dock 100
wuhan energy 100
wuhan finance 100
输出：
wuhan app
wuhan appdept
wuhan cpdept
wuhan dock
wuhan energy

按照规则：shenzhen人数最多的5个部门依次为：dbg abg cbg ebg zbg
输出规则：优先按照城市字典排序，先输出beijing 再shenzhen，部门按照字典升序输出
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type department struct {
	city           string
	departmentName string
	personNum      int64
}

type ByCtiy []department

func (a ByCtiy) Len() int {
	return len(a)
}

func (a ByCtiy) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByCtiy) Less(i, j int) bool {
	return a[i].city < a[j].city
}

func getTop5Department(departments []department) []department {
	ans := make([]department, 0)
	sort.Slice(departments, func(i, j int) bool {
		if departments[i].city < departments[j].city {
			return true
		}
		if departments[i].city == departments[j].city && departments[i].personNum > departments[j].personNum {
			return true
		}
		return false
	})
	//此处做取每个地区人数最多的5个部门，定义两个flag，一个当标志记录同部门人数最多的5个部门的标志为，一个标志该地址最后一个部门的下标
	l, r := 0, 0
	var flag = 0
	for range departments {
		if r < len(departments)-1 && departments[r].city == departments[r+1].city {
			if l-flag < 4 {
				l++ //同一个城市的部门不足五个的时候会一直数，数够5个右标会继续计数，直到该城市的走完
			}
			r++
		} else {
			if l-flag == 0 {
				ans = append(ans, departments[r:r+1]...)
				r++
			}
			if l-flag <= 4 && l-flag > 0 {
				ans = append(ans, departments[flag:l+1]...)
				r++
			}
			l, flag = r, r
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		if ans[i].city < ans[j].city {
			return true
		}
		if ans[i].city == ans[j].city && ans[i].departmentName < ans[j].departmentName {
			return true
		}
		return false
	})

	return ans
}

//func main() {
//	var n int
//	if _, err := fmt.Scanf("%d\n", &n); err != nil {
//		return
//	}
//
//	departments := make([]department, n)
//	reader := bufio.NewReader(os.Stdin)
//	for i := 0; i < n; i++ {
//		var err error
//		departments[i], err = readline(reader)
//		if err != nil {
//			return
//		}
//	}
//	ans := getTop5Department(departments)
//	fmt.Println("output:")
//	for _, v := range ans {
//		fmt.Println(v.city, v.departmentName)
//	}
//}

func readline(reader *bufio.Reader) (department, error) {
	lineBuf, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return department{}, fmt.Errorf(err.Error())
	}

	lineBuf = strings.TrimRight(lineBuf, "\r\n")
	line := strings.Split(lineBuf, " ")

	num, err := strconv.ParseInt(line[2], 10, 32)
	if err != nil {
		return department{}, fmt.Errorf(err.Error())
	}

	return department{line[0], line[1], num}, nil
}
