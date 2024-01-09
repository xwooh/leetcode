// Created by Yuan at 2024/01/09 10:43
// leetgo: dev
// https://leetcode.cn/problems/merge-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 先把 nums1 的数都放到后面空位上，然后再开始走合并操作

	// nums2 无值
	if n == 0 {
		return
	}

	// nums1 无值
	if m == 0 {
		for idx := range nums2 {
			nums1[idx] = nums2[idx]
		}
		return
	}

	// nums1 和 nums2 都有值
	for i := 0; i < m; i++ {
		nums1[len(nums1)-1-i] = nums1[m-1-i]
	}

	l, r := len(nums1)-m, 0
	p := 0
	for l < len(nums1) && r < n {
		if nums1[l] < nums2[r] {
			nums1[p] = nums1[l]
			l++
		} else {
			nums1[p] = nums2[r]
			r++
		}
		p++
	}

	if l == len(nums1) {
		for i := r; i < n; i++ {
			nums1[p] = nums2[i]
			p++
		}
	}

	if r == n {
		for i := l; i < len(nums1); i++ {
			nums1[p] = nums1[i]
			p++
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	m := Deserialize[int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	merge(nums1, m, nums2, n)
	ans := nums1

	fmt.Println("\noutput:", Serialize(ans))
}
