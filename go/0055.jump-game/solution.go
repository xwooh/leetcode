// Created by Yuan at 2024/01/19 15:09
// leetgo: dev
// https://leetcode.cn/problems/jump-game/

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

func canJump(nums []int) bool {
	// 最差也就一步一步往前走
	// 那就每前进一步看看最远的距离能到哪
	// 如果前进路上发现能到的最远距离比当前还要小，那就表示无法到达
	mx := 0
	for idx, v := range nums {
		// 当前能到的最远距离
		mx = max(mx, idx+v)
		if mx <= idx && idx < len(nums)-1 {
			// 能到的最远距离比当前还近，那就走不了了
			return false
		}
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := canJump(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
