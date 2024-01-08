// Created by Yuan at 2023/09/05 13:48
// leetgo: dev
// https://leetcode.cn/problems/form-smallest-number-from-two-digit-arrays/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func minx(nums []int) int {
	m := nums[0]
	for _, n := range nums[1:] {
		if n < m {
			m = n
		}
	}

	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minNumber(nums1 []int, nums2 []int) (ans int) {
	/*
		- 如果两个数组中存在相同的数字，取最小的相同数
		- 如果两个数组中不存在相同的数字，各取最小的数
	*/

	m := make(map[int]int, len(nums1))
	for _, n := range nums1 {
		m[n] = 1
	}

	// 两数组是否有相同的数
	s := 10
	for _, n := range nums2 {
		if _, ok := m[n]; ok {
			s = min(s, n)
		}
	}
	if s != 10 {
		ans = s
		return
	}

	// 找出两数组各自最小值
	m1 := minx(nums1)
	m2 := minx(nums2)

	if m1 < m2 {
		ans = 10*m1 + m2
	} else {
		ans = 10*m2 + m1
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	ans := minNumber(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
