// Created by Yuan at 2024/02/23 11:50
// leetgo: dev
// https://leetcode.cn/problems/JEj789/

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

func minCost(costs [][]int) (ans int) {
	// m[i] 表示第 i 阶段选择 红色、蓝色、绿色 这三种颜色分别对应的花费

	// 可知，初始第一阶段，这三种颜色的花费如下
	m := [...]int{costs[0][0], costs[0][1], costs[0][2]}
	ans = min(min(m[0], m[1]), m[2])

	for i := 1; i < len(costs); i++ {
		// 前一轮各情况成本
		m0 := m[0]
		m1 := m[1]
		m2 := m[2]

		// 当前各情况各成本
		m[0] = min(m1, m2) + costs[i][0]
		m[1] = min(m0, m2) + costs[i][1]
		m[2] = min(m0, m1) + costs[i][2]

		ans = min(min(m[0], m[1]), m[2])
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	costs := Deserialize[[][]int](ReadLine(stdin))
	ans := minCost(costs)

	fmt.Println("\noutput:", Serialize(ans))
}
