// Created by Yuan at 2023/11/06 10:35
// leetgo: dev
// https://leetcode.cn/problems/longest-increasing-subsequence/

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
func dp(nums []int, i int, table map[int]int) int {
	// 以索引 i 元素结尾的，最长子序列长度 l

	m := 1
	for j := i - 1; j >= 0; j-- {
		var dpn int
		if v, ok := table[j]; ok {
			dpn = v
		} else {
			dpn = dp(nums, j, table)
		}
		if nums[i] > nums[j] {
			m = max(m, dpn+1)
		}
	}
	table[i] = m
	return m
}

func lengthOfLIS(nums []int) (ans int) {
	table := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		ans = max(ans, dp(nums, i, table))
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := lengthOfLIS(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
