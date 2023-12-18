// Created by Yuan at 2023/10/13 18:48
// leetgo: dev
// https://leetcode.cn/problems/median-of-two-sorted-arrays/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getBorder(nums1, nums2 []int) (int, int) {
	if len(nums1) == 0 {
		return nums2[0], nums2[len(nums2)-1]
	} else if len(nums2) == 0 {
		return nums1[0], nums1[len(nums1)-1]
	}
	return min(nums1[0], nums2[0]), max(nums1[len(nums1)-1], nums2[len(nums2)-1])
}

func countOnce(nums []int, x int) int {
	// nums1, nums2 都是非递减数组
	// nums1, nums2 中 <= x 的总和，即: 二分找 <= x 的右边界

	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l+1)/2
		if nums[m] <= x {
			// 有点靠左，右移
			l = m
		} else {
			// 太靠右了，左移
			r = m - 1
		}
	}

	if nums[l] <= x {
		return l
	}

	return -1
}

func count(nums1, nums2 []int, x int) int {
	// nums1, nums2 都是非递减数组
	// nums1, nums2 中 <= x 的总和，即: 二分找 <= x 的右边界

	c1 := countOnce(nums1, x) + 1
	c2 := countOnce(nums2, x) + 1

	return c1 + c2
}

func findTopK(nums1 []int, nums2 []int, l, r, k int) int {
	// count(nums1, nums2, x) <= k 的左边界
	for l < r {
		x := l + (r-l)/2
		t := count(nums1, nums2, x)
		if t >= k {
			// 说明 x 太大了，左移
			r = x
		} else {
			// 说明 x 过小，右移
			l = x + 1
		}
	}

	if count(nums1, nums2, l) >= k {
		return l
	}
	return -1
}

func findMedianSortedArrays(nums1 []int, nums2 []int) (ans float64) {
	// l, r = min(nums1, nums2), max(nums1, nums2)
	// count(x) -> int: 返回 nums1, nums2 中 <= x 的个数 (第 k 小)
	// 如果奇数: k = ( len(nums1) + len(nums2) + 1 ) / 2
	// 如果偶数: k = ( len(nums1) + len(nums2) ) / 2

	l, r := getBorder(nums1, nums2)

	if (len(nums1)+len(nums2))%2 == 0 {
		// 偶数
		k := (len(nums1) + len(nums2)) / 2
		x1 := findTopK(nums1, nums2, l, r, k)
		x2 := findTopK(nums1, nums2, l, r, k+1)
		ans = (float64(x1) + float64(x2)) / 2.0
	} else {
		// 奇数
		k := (len(nums1) + len(nums2) + 1) / 2
		x := findTopK(nums1, nums2, l, r, k)
		ans = float64(x)
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	ans := findMedianSortedArrays(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
