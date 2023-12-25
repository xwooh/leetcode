// Created by Yuan at 2023/12/25 14:28
// leetgo: dev
// https://leetcode.cn/problems/daily-temperatures/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type T struct {
	// 当前温度
	t int
	// 比前面多少位高（自己也算一位）
	c int
	// 当前温度是哪一天的
	idx int
}

func dailyTemperatures(temperatures []int) (ans []int) {
	if len(temperatures) == 0 {
		return
	} else if len(temperatures) <= 1 {
		ans = []int{0}
		return
	}

	// 初始化
	ans = make([]int, len(temperatures))

	stack := []T{}

	for idx, v := range temperatures {
		// 当前温度
		e := T{v, 1, idx}

		for len(stack) > 0 && e.t > stack[len(stack)-1].t {
			// 当前温度比栈顶温度高
			// 栈顶元素出栈，比栈顶元素温度高的天数 = e.c
			// 出栈后栈顶元素 T.c 加给当前元素的 e.c
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top.idx] = e.c
			e.c += top.c
		}

		stack = append(stack, e)
	}

	// 栈里面剩下的是非递增数组，也就是说每一天都不存在后面比当天温度高的日子
	for _, s := range stack {
		ans[s.idx] = 0
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	temperatures := Deserialize[[]int](ReadLine(stdin))
	ans := dailyTemperatures(temperatures)

	fmt.Println("\noutput:", Serialize(ans))
}
