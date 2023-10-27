// Created by Yuan at 2023/10/27 11:31
// leetgo: dev
// https://leetcode.cn/problems/max-consecutive-ones-iii/

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

func longestOnes(nums []int, k int) (ans int) {

	var l, r int
	var zeroN int

	for r < len(nums) {
		if nums[r] == 0 {
			// 如果 nums[r] == 0 需要判断能否加进来，这时就要动左侧指针
			// 当前是 0，所以必须弹一个 0 出来
			for zeroN >= k {
				if nums[l] == 0 {
					// 如果弹出来的是 0，累计的 0 量 -1
					zeroN--
				}
				// 左移一位
				l++
			}
			zeroN++
		}
		ans = max(ans, r-l+1)
		// 左侧指针已固定，动右侧指针
		r++
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := longestOnes(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
