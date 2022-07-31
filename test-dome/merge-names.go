/*Implement the uniqueNames function. When passed two slices of names,
 it will return a slice containing the names that appear in either or both slices.
 The returned slice should have no duplicates.
For example, calling uniqueNames([]string{"Ava", "Emma", "Olivia"},
[]string{"Olivia", "Sophia", "Emma"}) should return a slice containing Ava, Emma, Olivia, and Sophia in any order.*/

package main

import (
	"fmt"
	"sort"
)

func uniqueNames(a, b []string) []string {
	dubliCheck := make(map[string]bool)
	unqNames := make([]string, 0)
	a = append(a, b...)
	sort.Strings(a)
	for _, val := range a {
		if !dubliCheck[val] {
			unqNames = append(unqNames, val)
			dubliCheck[val] = true
		}
	}
	return unqNames
}
func main() {
	sl1 := []string{"Ava", "Emma", "Olivia"}
	sl2 := []string{"Olivia", "Sophia", "Emma"}

	fmt.Println(uniqueNames(sl1, sl2))
}
