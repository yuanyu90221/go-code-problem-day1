package main

import (
	"fmt"

	"github.com/yuanyu90221/go-code-problem-day1/sol"
)

func main() {
	nums := []int{1, 3, 3, 3, 3}
	target := 3
	result := sol.SearchInsertionPosition(nums, target)
	fmt.Println(result)
}
