// Created by Yuan at 2023/10/30 22:17
// leetgo: dev
// https://leetcode.cn/problems/steps-to-make-array-non-decreasing/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func totalSteps(nums []int) (ans int) {
	stack := []int{}

	var M int
	for _, v := range nums {
		M = max(M, v)
	}

	var c int
	var poped bool
	for i := len(nums) - 1; i >= 0; i-- {
		for (len(stack) > 0 && nums[i] > stack[len(stack)-1]) || (poped && len(stack) == 1 && nums[i] == stack[len(stack)-1] && nums[i] < M) {
			stack = stack[:len(stack)-1]
			c++
			ans = max(ans, c)
			poped = true
		}
		poped = false
		// 清零，重新计算
		c = 0
		stack = append(stack, nums[i])
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := totalSteps(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
