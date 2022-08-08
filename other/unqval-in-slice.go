package main

import "fmt"

var slice = []int{1, 1, 2, 2, 3, 3, 4, 5, 5}

func main() {
	res := make(map[int]int)
	for _, val := range slice {
		v, ok := res[val]
		if ok {
			v = v + 1
			res[val] = v
		} else {
			res[val] = 1
		}
	}
	for key, val := range res {
		if val == 1 {
			fmt.Println("Unq value in slice is :", key)
			return
		}
	}
}
