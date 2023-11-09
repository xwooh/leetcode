// Created by Yuan at 2023/11/09 22:14
// leetgo: dev
// https://leetcode.cn/problems/greatest-sum-divisible-by-three/

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

func maxSumDivThree(nums []int) (ans int) {
	// 每次都用当前值去加上之前余为 0、1、2 的和
	// 然后计算所得的和应该属于余为 0、1、2 的哪个，并保持当前对应余的和都是最大值

	var pre0, pre1, pre2 int
	var m0, m1, m2 int

	l := len(nums)
	for i := 0; i < l; i++ {
		if (pre0+nums[i])%3 == 0 {
			m0 = max(m0, pre0+nums[i])
		} else if (pre1+nums[i])%3 == 0 {
			m0 = max(m0, pre1+nums[i])
		} else if (pre2+nums[i])%3 == 0 {
			m0 = max(m0, pre2+nums[i])
		}

		if (pre0+nums[i])%3 == 1 {
			m1 = max(m1, pre0+nums[i])
		} else if (pre1+nums[i])%3 == 1 {
			m1 = max(m1, pre1+nums[i])
		} else if (pre2+nums[i])%3 == 1 {
			m1 = max(m1, pre2+nums[i])
		}

		if (pre0+nums[i])%3 == 2 {
			m2 = max(m2, pre0+nums[i])
		} else if (pre1+nums[i])%3 == 2 {
			m2 = max(m2, pre1+nums[i])
		} else if (pre2+nums[i])%3 == 2 {
			m2 = max(m2, pre2+nums[i])
		}

		pre0 = m0
		pre1 = m1
		pre2 = m2

	}

	ans = pre0

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maxSumDivThree(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
