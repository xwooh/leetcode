// Created by Yuan at 2023/10/27 17:30
// leetgo: dev
// https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func sum(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}

func numOfSubarrays(arr []int, k int, threshold int) (ans int) {
	var l, r int

	for r < len(arr) {
		// 什么情况下左边界要移动
		// 1. 长度 >= k
		// 2. sum >= k * threshold
		for (r - l + 1) >= k {
			if sum(arr[l:r+1]) >= k*threshold {
				ans++
			}
			l++
		}
		r++
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	arr := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	threshold := Deserialize[int](ReadLine(stdin))
	ans := numOfSubarrays(arr, k, threshold)

	fmt.Println("\noutput:", Serialize(ans))
}
