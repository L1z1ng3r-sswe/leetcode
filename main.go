package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var result [][]int
	fmt.Println("Starting permutation generation")
	backtrack(nums, 0, &result)
	fmt.Println("Finished permutation generation")
	return result
}

func backtrack(nums []int, start int, result *[][]int) {
	fmt.Printf("Entering backtrack: nums = %v, start = %d\n", nums, start)
	if start == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		fmt.Printf("Complete permutation found: %v\n", temp)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		fmt.Printf("Swapping: nums[%d] = %d, nums[%d] = %d\n", start, nums[start], i, nums[i])
		nums[start], nums[i] = nums[i], nums[start]

		fmt.Printf("Recursing with nums = %v and start = %d\n", nums, start+1)
		backtrack(nums, start+1, result)

		fmt.Printf("Backtracking: Re-swapping nums[%d] = %d, nums[%d] = %d\n", start, nums[start], i, nums[i])
		nums[start], nums[i] = nums[i], nums[start]
	}
	fmt.Printf("Exiting backtrack: nums = %v, start = %d\n", nums, start)
}

func main() {
	nums := []int{1, 2, 3}
	permutations := permute(nums)
	fmt.Println("Generated permutations:")
	for _, perm := range permutations {
		fmt.Println(perm)
	}
}
