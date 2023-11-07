// Created by Yuan at 2023/11/07 18:28
// leetgo: dev
// https://leetcode.cn/problems/min-cost-climbing-stairs/

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

func minCostClimbingStairs(cost []int) (ans int) {
	// dp(i) 表示走 i 位置的最小花费
	// 当前  = min(上一步到, 上上步到) + 当前cost
	// dp(i) = min(dp(i-1), dp(i-2)) + cost[i]

	table := make(map[int]int)

	l := len(cost)

	for i := 0; i <= l; i++ {
		if i == 0 {
			table[i] = cost[0]
		} else if i == 1 {
			table[i] = cost[1]
		} else {
			var c int
			if i == l {
				c = 0
			} else {
				c = cost[i]
			}
			table[i] = min(table[i-1], table[i-2]) + c
		}
	}
	ans = table[l]
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	cost := Deserialize[[]int](ReadLine(stdin))
	ans := minCostClimbingStairs(cost)

	fmt.Println("\noutput:", Serialize(ans))
}
