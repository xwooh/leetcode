// Created by Yuan at 2023/12/27 15:22
// leetgo: dev
// https://leetcode.cn/problems/house-robber/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) (ans int) {
	// 第 i 家偷/不偷的最大金额
	if len(nums) == 0 {
		return
	} else if len(nums) == 1 {
		ans = nums[0]
		return
	} else if len(nums) == 2 {
		ans = max(nums[0], nums[1])
		return
	}

	// 上一次和上上次偷取的最大金额
	ppre := nums[0]
	pre := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		// 本次最大偷取金额 = max(上次最大(本次不偷), 上上次最大+本次偷)
		ans = max(pre, ppre+nums[i])

		ppre = pre
		pre = ans
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := rob(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
