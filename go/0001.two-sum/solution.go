// Created by Yuan at 2023/10/13 13:40
// leetgo: dev
// https://leetcode.cn/problems/two-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func twoSum(nums []int, target int) (ans []int) {
	m := make(map[int]int)

	for idx, v := range nums {
		ret := target - v
		if midx, ok := m[ret]; ok {
			return []int{midx, idx}
		} else {
			// 记录当前这个数的位置
			m[v] = idx
		}
	}

	return []int{-1, -1}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := twoSum(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
