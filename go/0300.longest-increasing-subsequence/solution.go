// Created by Yuan at 2023/11/06 10:35
// leetgo: dev
// https://leetcode.cn/problems/longest-increasing-subsequence/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func bsl(nums []int, t int) int {
	l, r := 0, len(nums)

	for l < r {
		m := l + (r-l)/2
		if nums[m] >= t {
			// 找左边界
			r = m
		} else {
			l = m + 1
		}
	}

	if nums[l] >= t {
		return l
	}

	return -1
}

func lengthOfLIS(nums []int) (ans int) {
	// 维护一个数组 p，p[i] 表示最长上升子序列长度为 i+1 的时候，所有符合这个长度序列中的最小结尾值。
	// 比如 [2, 4, 3, 5] 长度为 2 的上升子序列有 [2, 3] [2, 4] [2, 5] [3, 5]，那对应 p[1] = min([3, 4, 5]) = 3
	p := []int{}

	for _, v := range nums {
		if len(p) == 0 {
			p = append(p, v)
		} else if v > p[len(p)-1] {
			p = append(p, v)
		} else {
			// 在 p 中寻找 >= v 的左边界（找到第一个不小于 v 的值），并替换掉
			pIdx := bsl(p, v)
			if pIdx != -1 {
				p[pIdx] = v
			}
		}
	}

	return len(p)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := lengthOfLIS(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
