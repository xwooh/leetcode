// Created by Yuan at 2023/11/30 12:10
// leetgo: dev
// https://leetcode.cn/problems/kth-largest-element-in-an-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findKthLargest(nums []int, k int) (ans int) {
	// 最小最大值
	s, b := nums[0], nums[0]
	for _, v := range nums[1:] {
		s = min(s, v)
		b = max(b, v)
	}

	// [s, b] 之间选一个数 x，使得所有的数 >= x (左边界)
	for s < b {
		x := s + (b-s)/2 + 1

		kc := 0
		for _, v := range nums {
			if v >= x {
				if kc == 0 {
					ans = v
				} else {
					ans = min(ans, v)
				}
				kc++
			}
		}
		if kc == k {
			return
		} else if kc < k {
			// 这个数大了，小一点 => 左移
			b = x - 1
		} else {
			// 这个数小了，大一点 => 右移
			s = x
		}

	}

	kc := 0
	for _, v := range nums {
		if v >= s {
			if kc == 0 {
				ans = v
			} else {
				ans = min(ans, v)
			}
			kc++
		}
	}
	if kc >= k {
		return
	}

	return -1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := findKthLargest(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
