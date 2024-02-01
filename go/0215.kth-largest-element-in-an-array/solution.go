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

func count(nums []int, v int) (int, int) {
	// returns:
	//   - nums 中 >=v 的个数
	//   - <=v 的最大值，因为要找的是 nums 中的值
	c := 0
	key := math.MaxInt
	for _, n := range nums {
		if n >= v {
			c++
			key = min(key, n)
		}
	}

	return c, key
}

func findKthLargest(nums []int, k int) (ans int) {
	// 最小最大值
	l, r := nums[0], nums[0]
	for _, v := range nums[1:] {
		l = min(l, v)
		r = max(r, v)
	}

	for l <= r {
		// 左边界就是变化的点，所以左边界必定在 nums 中
		m := l + (r-l)/2

		c, kv := count(nums, m)

		if c == k {
			// 求左边界，那就继续往左偏移
			return kv
		} else if c > k {
			// >= m 的数多了，也就是 m 有点小需要变大，那就往右偏移
			// 这里很神奇，如果 nums 中大值比较少，相同小值又很多，
			//	那就会往右偏，但是一旦右偏，c 肯定 < k，
			//		但巧就巧在 c < k 条件下，r 左移后的值就是想要的 ans
			//		所以最后 return r
			l = m + 1
		} else {
			// >= m 的数少了，也就是 m 有点大需要变小，那就往左偏移
			r = m - 1
		}
	}

	return r
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := findKthLargest(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
