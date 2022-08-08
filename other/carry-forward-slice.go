/*

case1:
	slice:=[1,2,3]
	expected:- [1,2,4]

case2:=
	slice:=[1,2,9]
	expected:- [1,3,0]
case3:=
	slice:=[9,9,9]
	expected:- [1,0,0,0]
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

var slice = []int{9, 9, 9}

func main() {
	temp := ""
	res := make([]int, 0)
	for _, val := range slice {
		a := strconv.Itoa(val)
		temp += a
	}
	v, _ := strconv.Atoi(temp)
	v += 1
	m := strconv.Itoa(v)
	for _, p := range strings.Split(m, "") {
		o, _ := strconv.Atoi(p)
		res = append(res, o)

	}
	fmt.Println(res)
}
