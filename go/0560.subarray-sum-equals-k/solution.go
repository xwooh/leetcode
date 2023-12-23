// Created by Yuan at 2023/12/22 11:05
// leetgo: dev
// https://leetcode.cn/problems/subarray-sum-equals-k/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

func n2(nums []int, k int) (ans int) {
	var l, r, s int

	for l < len(nums) {
		s += nums[r]
		if s == k {
			ans++
			// 后续可能还能满足，走完所有的
		}
		r++

		if r >= len(nums) {
			// 走完了都没
			// 重置
			s = 0
			l++
			r = l
		}
	}

	return
}

// @lc code=begin

func subarraySum(nums []int, k int) (ans int) {
	// m 是前缀和出现的次数
	m := map[int]int{0: 1}

	var sum int
	for _, v := range nums {
		sum += v

		if c, ok := m[sum-k]; ok {
			// k = preSum[current] - preSum[somePre]
			// 只要之前出现过多次 preSum[current] - k 个前缀和，那就表示有这么多次机会
			ans += c
		}

		// 把当前前缀和记上
		m[sum]++

	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := subarraySum(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
