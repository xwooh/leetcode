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
	// m[i] 表示余为 0 1 2 时的最大元素和

	m := [3]int{}

	for i := 0; i < len(nums); i++ {
		// 之前余为 0 1 2 的三种情况，加上当前元素 nums[i] 后和的三种情况
		s0 := m[0] + nums[i]
		s1 := m[1] + nums[i]
		s2 := m[2] + nums[i]

		// 三种加上后和的三种余数情况
		// s0 s1 s2 模上 3 之后的余数肯定还是 0 1 2，所以 +nums[i] 后的三种情况就是下面这三种
		m[s0%3] = max(m[s0%3], s0)
		m[s1%3] = max(m[s1%3], s1)
		m[s2%3] = max(m[s2%3], s2)
	}

	ans = m[0]

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maxSumDivThree(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
