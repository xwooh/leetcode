// Created by Yuan at 2024/03/11 21:59
// leetgo: dev
// https://leetcode.cn/problems/circular-array-loop/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func walk(current, step int, n int) int {
	next := current + step%n

	if next >= n {
		next = next - n
	} else if next < 0 {
		next = n + next
	}

	return next
}

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		s := i
		f := walk(s, nums[s], n)
		for nums[s]*nums[f] > 0 && nums[walk(f, nums[f], n)]*nums[s] > 0 {
			if s == f {
				if s == walk(s, nums[s], n) {
					// 确保 k > 1
					break
				}
				return true
			}
			s = walk(s, nums[s], n)
			f = walk(f, nums[f], n)
			f = walk(f, nums[f], n)
		}
	}

	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := circularArrayLoop(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
