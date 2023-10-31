// Created by Yuan at 2023/10/31 20:50
// leetgo: dev
// https://leetcode.cn/problems/maximum-subarray/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func max(nums []int) (ans int) {
	ans = nums[0]
	for _, v := range nums[1:] {
		if v > ans {
			ans = v
		}
	}
	return
}

func maxSubArray(nums []int) (ans int) {
	// dp[i] 表示 nums[i] 为结尾的最大连续数组和

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else {
			if dp[i-1] <= 0 {
				dp[i] = nums[i]
			} else {
				dp[i] = dp[i-1] + nums[i]
			}
		}
	}

	ans = max(dp)

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maxSubArray(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
