package main

import (
	"strconv"
	"strings"
)

func buildToArr(input string) []int {
	input = strings.Trim(input, "[")
	input = strings.Trim(input, "]")
	strs := strings.Split(input, ",")
	rs := make([]int, len(strs))

	for i, str := range strs {
		rs[i], _ = strconv.Atoi(str)
	}
	return rs
}
