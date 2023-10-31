// Created by Yuan at 2023/10/31 22:17
// leetgo: dev
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

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

func maxProfit(prices []int) (ans int) {

	if len(prices) == 1 {
		return
	}

	stack := []int{prices[0]}

	for i := 1; i < len(prices); i++ {
		for len(stack) > 0 && prices[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, prices[i])

		ans = max(prices[i]-stack[0], ans)
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	prices := Deserialize[[]int](ReadLine(stdin))
	ans := maxProfit(prices)

	fmt.Println("\noutput:", Serialize(ans))
}
