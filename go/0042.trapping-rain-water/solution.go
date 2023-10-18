// Created by Yuan at 2023/10/18 22:48
// leetgo: dev
// https://leetcode.cn/problems/trapping-rain-water/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func trap(height []int) (ans int) {
	// 双指针，分别记录左右两边当前的最大值
	// 左右指针移动的时候，当前指针所在柱形左右最高点是知道的（记为 maxL, maxR）
	//	如果 maxR > maxL，那么此时可算左指针所在的柱子能接的雨水 = maxL - height[l]
	//	反之 maxL > maxR，可算得右指针所在的柱子能接的雨水 = maxR - height[r]

	l, r := 0, len(height)-1
	var maxL, maxR int
	for l < r {
		maxL = max(maxL, height[l])
		maxR = max(maxR, height[r])

		if maxR > maxL {
			ans += maxL - height[l]
			l++
		} else {
			ans += maxR - height[r]
			r--
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	height := Deserialize[[]int](ReadLine(stdin))
	ans := trap(height)

	fmt.Println("\noutput:", Serialize(ans))
}
