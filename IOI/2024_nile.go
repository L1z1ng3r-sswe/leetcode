package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(calculateCosts([]int{15, 12, 2, 10, 21}, []int{5, 4, 5, 6, 3}, []int{1, 2, 2, 3, 2}, []int{5, 9, 1}))
}

func calculateCosts(w []int, a []int, b []int, e []int) []int {
	if len(w) != len(a) || len(a) != len(b) {
		return []int{}
	}

	var N = len(w)

	var combinedW []struct{ num, aCost, bCost int }
	for i := range w {
		combinedW = append(combinedW, struct{ num, aCost, bCost int }{num: w[i], aCost: a[i], bCost: b[i]})
	}

	sort.Slice(combinedW, func(i, j int) bool {
		return combinedW[i].num < combinedW[j].num
	})

	fmt.Println(combinedW)

	var res []int

	for _, d := range e {
		var count int

		left := 0
		right := N - 1

		for left != right {
			if combinedW[left].num-d < combinedW[right].num {
				count += min(combinedW[left].bCost, combinedW[right].bCost)
				left++
			} else {
				right++
			}
		}
	}

	return w
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
