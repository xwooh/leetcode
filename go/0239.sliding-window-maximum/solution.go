// Created by Yuan at 2024/01/08 16:46
// leetgo: dev
// https://leetcode.cn/problems/sliding-window-maximum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxOne(nums []int) (int, int) {
	idx := 0
	m := nums[0]
	for i, v := range nums[1:] {
		if v > m {
			m = v
			idx = i
		}
	}
	return m, idx
}

func add(wis *[]int, idx int, nums []int) {
	for len(*wis) > 0 && nums[idx] >= nums[(*wis)[len(*wis)-1]] {
		// 当前元素比栈顶元素大，弹出栈顶元素
		*wis = (*wis)[:len(*wis)-1]
	}

	*wis = append(*wis, idx)
}

func maxSlidingWindow(nums []int, k int) (ans []int) {
	// 滑动窗口大小为 1 或者数组大小为 1
	if k == 1 || len(nums) == 1 {
		ans = nums
		return
	}

	// 滑动窗口比 nums 还大，那就是最大值
	if k >= len(nums) {
		m, _ := MaxOne(nums)
		ans = []int{m}
		return
	}

	// 维护单调队列
	// 为了方便判断最大值是否还在当前滑动窗口里面，
	//	单调队列里面存储的是元素索引

	// 初始化
	wis := []int{}
	for i := 0; i < k; i++ {
		add(&wis, i, nums)
	}
	ans = append(ans, nums[wis[0]])

	for i := range nums[k:] {
		idx := i + k

		// 确保队头元素在当前滑动窗口内
		for len(wis) > 0 && wis[0] <= idx-k {
			// pop 掉队头元素
			wis = wis[1:]
		}
		add(&wis, idx, nums)

		ans = append(ans, nums[wis[0]])
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := maxSlidingWindow(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
