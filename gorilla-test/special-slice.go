package main

import (
	"fmt"
	"math"
	"sync"
)

var m = []int{2, 2, 2, 2, 8}
var result []int

func main() {
	fmt.Println(SpecialArray(0, m))
}

func SpecialArray(n int, arr []int) []int {
	var result []int

	for _, val := range arr {
		var L, R int
		firCount := 0
		if val == 1 {
			fmt.Println("invalid slice")
			return nil
		}
		for i := 2; i <= int(math.Floor(float64(val)/2)); i++ {
			if val%i == 0 {
				firCount++
				break
			}
		}
		if firCount == 0 && val != 1 {
			bat := (val + val) * val
			result = append(result, bat)
		} else {
			temp1, temp2 := val, val
			var wg sync.WaitGroup
			wg.Add(2)
			go func(w *sync.WaitGroup) {
				for L == 0 {
					var isPrime bool

					temp1++

					for j := 2; j <= temp1/2; j++ {
						if temp1%j == 0 {
							isPrime = true
						}
					}
					if !isPrime && temp1 != 1 {
						L = temp1
						w.Done()

					}
				}

			}(&wg)
			go func(w *sync.WaitGroup) {
				for R == 0 {
					var isPrime bool
					temp2--

					for j := 2; j <= temp2/2; j++ {
						if temp2%j == 0 {
							isPrime = true
						}
					}
					if !isPrime && temp2 != 1 {
						R = temp2
						w.Done()

					}
				}

			}(&wg)
			wg.Wait()

			bat := (L + R) * val
			result = append(result, bat)
		}
	}
	return result
}
