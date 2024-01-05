// Created by Yuan at 2024/01/05 14:55
// leetgo: dev
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxProfit(prices []int) (ans int) {
	// 压入更小的数
	// 一旦有新的数比栈顶元素大，弹出栈顶元素，压入新的大数
	// 如果新的数比栈顶元素小，那就弹出栈顶的两个数，计算这两个数之差。

	stack := []int{}

	for _, p := range prices {
		for len(stack) >= 1 {
			if p <= stack[len(stack)-1] {
				if len(stack) == 1 {
					// 弹出栈顶小的数
					stack = stack[:len(stack)-1]
				} else {
					// 弹出栈顶的两个元素
					one := stack[len(stack)-1]
					two := stack[len(stack)-2]
					stack = stack[:len(stack)-2]
					ans += (one - two)
				}
			} else {
				// 当前元素比栈顶大
				// 弹出栈顶元素
				if len(stack) > 1 {
					stack = stack[:len(stack)-1]
				}
				break
			}
		}
		// 当前数比栈顶大，直接压入
		stack = append(stack, p)
	}

	if len(stack) >= 2 {
		ans += (stack[len(stack)-1] - stack[0])
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
