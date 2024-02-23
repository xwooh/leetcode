// Created by Yuan at 2024/01/04 14:40
// leetgo: dev
// https://leetcode.cn/problems/coin-change-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func change(amount int, coins []int) (ans int) {
	// f[j] = f[j]  // 不选 i
	//      = f[j-Ci]  // 选 i
	// 两种方案之和，即：f[j] += f[j-Ci]
	f := make([]int, amount+1)
	f[0] = 1

	for i := 0; i < len(coins); i++ {
		// j < c 的时候无法考虑 c 这个物品，直接从 c 开始考虑
		for j := coins[i]; j <= amount; j++ {
			// 选
			f[j] += f[j-coins[i]]
		}
	}

	ans = f[amount]
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	amount := Deserialize[int](ReadLine(stdin))
	coins := Deserialize[[]int](ReadLine(stdin))
	ans := change(amount, coins)

	fmt.Println("\noutput:", Serialize(ans))
}
