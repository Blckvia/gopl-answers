// ex5.15 create min, max, sum functions using variadic patter
package main

import "fmt"

func main() {
	fmt.Printf("Max value is: %d\n", max(6, 2, 3, 4))
	fmt.Printf("Min value is: %d\n", min(6, 2, 3, 4))
	fmt.Printf("Sum is: %d\n", sum(6, 2, 3, 4))
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}

func sum(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	for _, num := range nums {
		result += num
	}

	return result
}
