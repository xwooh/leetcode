// Created by Yuan at 2023/11/30 12:10
// leetgo: dev
// https://leetcode.cn/problems/kth-largest-element-in-an-array/

package main

import (
	"bufio"
	"fmt"
	"math"
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

func count(nums []int, p int) (ans int, x int) {
	x = math.MinInt
	for _, v := range nums {
		if v >= p {
			ans++
			x = max(x, v)
		}
	}
	return
}

func findKthLargest(nums []int, k int) (ans int) {
	// 值域二分

	// 先拿到值域的范围
	l := math.MaxInt
	r := math.MinInt
	for _, v := range nums {
		l = min(l, v)
		r = max(r, v)
	}

	// 在值域内二分，求左边界
	//	x 取的是符合条件的右边界值
	var x int
	for l < r {
		x = l + (r-l+1)/2
		c, x := count(nums, x)
		if c == k {
			r = x
		} else if c > k {
			// val 太小了，导致 c 过大
			l = x + 1
		} else {
			// val 太大了，导致 c 过小
			r = x - 1
		}
	}

	xc := 0
	for _, v := range nums {
		if x >= v {
			xc++
		}
	}

	if xc < k {
		return -1
	}

	return x
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := findKthLargest(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
