// Created by Yuan at 2023/11/28 18:51
// leetgo: dev
// https://leetcode.cn/problems/jump-game-ii/

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

func jump(nums []int) (ans int) {
	if len(nums) <= 1 {
		// 只有一个，那就不用跳了，已经到终点了
		return
	}

	m := 0
	for idx := 0; idx < len(nums); {
		v := nums[idx]

		// 看看当前位置能到的最远距离
		// 下一跳的最远距离
		m = max(m, idx+v)

		// 要想到 m，那肯定得跳一下
		ans++

		if m >= len(nums)-1 {
			// 下一跳就能完成了
			return
		}

		// 下一跳还不能完成
		// 在 [idx, idx+v] 中取下下一跳能到的最远距离
		np := idx + 1
		for i := 2; i <= v; i++ {
			p := idx + i
			if nums[p]+p >= nums[np]+np {
				// 如果这个位置能到的最远距离更远，就选这个
				np = p
			}
		}
		idx = np
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := jump(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
