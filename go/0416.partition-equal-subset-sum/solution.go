// Created by Yuan at 2024/02/23 18:31
// leetgo: dev
// https://leetcode.cn/problems/partition-equal-subset-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func sum(nums []int) (ans int) {
	for _, v := range nums {
		ans += v
	}
	return
}

func bagCount(nums []int, v int) bool {
	// 01背包正好凑 v 是否可行
	// f[j] 表示正好凑满空间 j 时的可行解数
	//       不考虑 i 物品 = f[j]
	// f[j] =
	//        考虑 i 物品 = f[j - Vi]

	f := make([]bool, v+1)
	f[0] = true // 要凑空间为 0 的背包，肯定是可以的
	for _, n := range nums {
		for j := v; j >= n; j-- {
			f[j] = f[j] || f[j-n]
		}
	}

	return f[v]
}

func canPartition(nums []int) bool {
	// 转成 01背包问题
	// 题意要分成两半相等的，那就看 sum / 2 容量的背包有无解

	s := sum(nums)
	if s%2 == 1 {
		return false
	}

	v := s / 2
	return bagCount(nums, v)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := canPartition(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
