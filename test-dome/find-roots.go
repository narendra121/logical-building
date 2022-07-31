/*
Implement the function findRoots to find the roots of the quadratic equation: ax2 + bx + c = 0.
If the equation has only one solution, the function should return that solution as both results.
The equation will always have at least one solution.

The roots of the quadratic equation can be found with the following formula: A quadratic equation.

For example, the roots of the equation 2x2 + 10x + 8 = 0 are -1 and -4.
*/

package main

import (
	"fmt"
	"math"
)

func findRoots(a, b, c float64) (float64, float64) {
	inRootVal := (b * b) - (4 * a * c)
	rootVal := math.Sqrt(inRootVal)
	upVal1 := -b - float64(rootVal)
	upVal2 := -b + float64(rootVal)

	divider := 2 * a
	return upVal1 / divider, upVal2 / divider
}

func main() {
	x1, x2 := findRoots(2, 10, 8)
	fmt.Printf("Roots: %f, %f", x1, x2)
}
