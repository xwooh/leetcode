// Created by Yuan at 2023/10/28 22:55
// leetgo: dev
// https://leetcode.cn/problems/next-greater-element-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func nextGreaterElements(nums []int) (ans []int) {
	if len(nums) == 1 {
		return []int{-1}
	}

	// 初始化 ans
	ans = make([]int, len(nums))

	var mxN int = nums[0]
	// 存放元素的索引
	stack := []int{}

	// 先过一遍单调栈
	for idx, v := range nums {
		// 当前元素要比栈头元素大，弹出栈头元素
		// 当前元素的索引 append 到 ans
		for len(stack) > 0 && v > nums[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
			mxN = max(mxN, v)
		}
		// 当前元素为栈中最大元素
		stack = append(stack, idx)
	}

	// 看看还有哪些索引位置还没找到下一个较大元素
	newStack := []int{}
	// 再用原数组过一遍单调栈
	for _, v := range nums {
		for len(stack) > 0 && (v > nums[stack[len(stack)-1]] || nums[stack[len(stack)-1]] == mxN) {
			if v > nums[stack[len(stack)-1]] {
				ans[stack[len(stack)-1]] = v
				newStack = append(newStack, v)
				stack = stack[:len(stack)-1]
			} else if nums[stack[len(stack)-1]] == mxN {
				ans[stack[len(stack)-1]] = -1
				newStack = append(newStack, -1)
				stack = stack[:len(stack)-1]
			}
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := nextGreaterElements(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
