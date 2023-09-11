// ex4.4 rotates a slice of ints by k positions to the left.
package rotate

import "fmt"

func rotate(nums []int, k int) []int {
	if k < 0 || len(nums) == 0 {
        return nums
    }
	r := len(nums) - k%len(nums)
    nums = append(nums[r:], nums[:r]...)

	return nums

}

func main() {
	fmt.Println(rotate([]int{1,2,3}, 1))
}