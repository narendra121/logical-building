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

import "fmt"

var slice = []int{0, 5, 9}

func main() {
	last := slice[len(slice)-1]
	last = last + 1
	first := slice[0]
	resp := make([]int, 0)

	if last > 9 {
		s := 1
		for i := len(slice) - 1; i >= 0; i-- {
			temp := 0
			if s > 0 {
				temp = slice[i] + 1
				if temp > 9 {
					s = 1
					slice[i] = 0
				} else {
					s = 0
					slice[i] = temp
				}
			}

		}
		if first == 9 {
			resp = append(resp, 1)
			resp = append(resp, slice...)

		} else {
			resp = append(resp, slice...)
		}
	} else {
		slice[len(slice)-1] = slice[len(slice)-1] + 1
		resp = append(resp, slice...)

	}
	fmt.Println(resp)
}
