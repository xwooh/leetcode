// Created by Yuan at 2024/01/03 14:50
// leetgo: dev
// https://leetcode.cn/problems/coin-change/

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

func coinChange(coins []int, amount int) (ans int) {
	// m[i] = x
	// 表示凑到 i 金额，所需的最小硬币数为 x
	// 那么凑到 j 金额，所需的最小硬币数为 min([m[j-c]+1 for c in coins])

	if amount < 0 {
		return -1
	} else if amount == 0 {
		return 0
	}

	// 初始化
	DEFAULT_V := amount + 1
	m := make([]int, DEFAULT_V)
	for idx := range m {
		m[idx] = DEFAULT_V
	}

	for x := 1; x <= amount; x++ {
		for _, c := range coins {
			if x == c {
				m[x] = min(m[x], 1)
			} else {
				w := x - c
				if w < 0 {
					// 无解
					continue
				} else {
					wc := m[w]
					if wc == DEFAULT_V {
						// 说明无解
						continue
					} else {
						m[x] = min(m[x], wc+1)
					}
				}
			}
		}
	}

	ans = m[amount]
	if ans == DEFAULT_V {
		ans = -1
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	coins := Deserialize[[]int](ReadLine(stdin))
	amount := Deserialize[int](ReadLine(stdin))
	ans := coinChange(coins, amount)

	fmt.Println("\noutput:", Serialize(ans))
}
