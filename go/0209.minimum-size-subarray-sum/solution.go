// Created by Yuan at 2023/10/26 13:06
// leetgo: dev
// https://leetcode.cn/problems/minimum-size-subarray-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func minAns(ans, count int) int {
	if ans == 0 {
		return count
	}
	if count < ans {
		return count
	}
	return ans
}

func minSubArrayLen(target int, nums []int) (ans int) {
	pipe := []int{}
	pSum := 0
	count := 0

	for _, v := range nums {
		// 缩小滑动窗口
		for len(pipe) > 0 && pSum >= target {
			// 此时还是满足的，下面 -1 后不一定满足
			ans = minAns(ans, count)
			// 弹左边
			lv := pipe[0]
			pipe = pipe[1:]
			pSum -= lv
			count--
		}
		pipe = append(pipe, v)
		pSum += v
		count++
	}

	if pSum < target {
		return
	}

	// 最后可能还可以继续缩小滑动窗口
	for len(pipe) > 0 && pSum >= target {
		ans = minAns(ans, count)
		lv := pipe[0]
		pipe = pipe[1:]
		pSum -= lv
		count--
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	target := Deserialize[int](ReadLine(stdin))
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := minSubArrayLen(target, nums)

	fmt.Println("\noutput:", Serialize(ans))
}
