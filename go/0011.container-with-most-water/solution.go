// Created by Yuan at 2023/11/10 11:50
// leetgo: dev
// https://leetcode.cn/problems/container-with-most-water/

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) (ans int) {

	// ans = (r - l) * min(height[r], height[l])

	l, r := 0, len(height)-1
	var plv, prv int

	for l < r {
		lv := height[l]
		rv := height[r]

		if lv <= plv {
			l++
			continue
		}

		if rv <= prv {
			r--
			continue
		}

		ans = max(ans, (r-l)*min(lv, rv))

		if lv > rv {
			r--
		} else {
			l++
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	height := Deserialize[[]int](ReadLine(stdin))
	ans := maxArea(height)

	fmt.Println("\noutput:", Serialize(ans))
}
