// Created by Yuan at 2024/02/24 19:19
// leetgo: dev
// https://leetcode.cn/problems/target-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func sum(nums []int) (ans int) {
	for _, n := range nums {
		ans += n
	}
	return
}

func findTargetSumWays(nums []int, target int) (ans int) {
	// 挑选其中某些物品，使得 sum - 2*(m) = target
	// 即选出背包体积为 m = (sum - target) / 2 的方案数
	s := sum(nums)
	if target > s {
		return
	}

	m2 := s - target
	if m2%2 != 0 {
		// 不整除，则表示无法凑出解
		return
	}

	m := m2 / 2
	f := make([]int, m+1)
	f[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := m; j >= nums[i]; j-- {
			f[j] += f[j-nums[i]]
		}
	}

	ans = f[m]
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := findTargetSumWays(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
