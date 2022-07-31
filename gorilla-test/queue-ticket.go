package main

import "fmt"

func LastPerson(n int, k int) int {
	queue := make([]int, 0)
	result := 0
	for person := 1; person <= n; person++ {
		queue = append(queue, person)
	}
	temp := 2

	for len(queue) > k {
		lgth := (len(queue))
		if temp%2 == 0 {
			queue = queue[k:]
		} else {
			queue = queue[:lgth-k]
		}
		temp++
	}

	result = queue[len(queue)-1]

	return result
}

func main() {

	fmt.Println(LastPerson(2, 1))
}
