// Created by Yuan at 2023/10/25 21:29
// leetgo: dev
// https://leetcode.cn/problems/132-pattern/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func find132pattern(nums []int) bool {
	// 132 模式，那么如果 2 是比 3 小的最大值
	// 这样找到的 3 和 2 能匹配到的 1 可能性最大
	// => 从右往左找

	if len(nums) < 3 {
		return false
	}

	stack := []int{}

	var n2 int
	var poped bool

	// 从最后一个开始遍历
	for r := len(nums) - 1; r >= 0; r-- {
		if poped && len(stack) != 0 {
			// n2 有了，如果 stack 不为空，那 n3 也有了
			// 此时要遍历的值就是 n1 备选项
			if nums[r] < n2 {
				return true
			}
		}
		for len(stack) != 0 && nums[r] > stack[len(stack)-1] {
			// 弹出比当前值小的数
			// 这样能找到当前栈中比当前数小的最大值
			// 然后把当前值入栈
			n2 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			poped = true
		}
		stack = append(stack, nums[r])
	}

	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := find132pattern(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
