// Created by Yuan at 2023/10/22 11:13
// leetgo: dev
// https://leetcode.cn/problems/majority-element/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func majorityElement(nums []int) (ans int) {
	// 摩尔投票：不一样的抵消

	// 题目已假设 nums 非空
	ans = nums[0]
	candCount := 0
	for _, v := range nums[1:] {
		if v != ans {
			if candCount > 0 {
				candCount--
			} else {
				// 题目假设必定存在多数元素，所以这边直接把众数认为 v 也没关系
				// 此时众数肯定不是原本的 ans 了，随便给个 v 就好了
				ans = v
				candCount = 0
			}
		} else {
			candCount++
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := majorityElement(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
