package sol

import "fmt"

func BinarySearchIndex(nums *[]int, target int, left int, right int) int {
	mid := (left + right) / 2
	l := left
	r := right
	fmt.Printf("mid = %d, l=%d, r=%d, target=%d\n", mid, l, r, target)
	if l == r {
		return l
	}
	if target > (*nums)[mid] { // target > mid
		return BinarySearchIndex(nums, target, mid+1, r)
	} else {
		return BinarySearchIndex(nums, target, l, mid)
	}
}
func SearchInsertionPosition(nums []int, target int) int {
	result := 0
	right := len(nums) - 1
	result = BinarySearchIndex(&nums, target, 0, right)
	return result
}
