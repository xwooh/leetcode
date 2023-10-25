// Created by Yuan at 2023/10/25 13:33
// leetgo: dev
// https://leetcode.cn/problems/longest-valid-parentheses/

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

func longestValidParentheses(s string) (ans int) {
	idx := make([]int, len(s))

	stack := []int{}
	for i, c := range s {
		if string(c) == "(" {
			// 左括号入栈，value 置为 1
			stack = append(stack, i)
			idx[i] = 1
		} else if string(c) == ")" {
			if len(stack) == 0 {
				// 没有左括号与之配对，直接丢弃，value 置为 1
				idx[i] = 1
			} else {
				// 左括号出栈，并把左括号的 value 置为 0
				l := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				idx[l] = 0
				// 右括号匹配上了，value 也置为 0
				idx[i] = 0
			}
		}
	}

	var length int
	for _, v := range idx {
		if v == 1 {
			// 碰壁，重新计数
			length = 0
		} else {
			// 该位置能匹配上的，长度 +1
			length++
			ans = max(ans, length)
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := longestValidParentheses(s)

	fmt.Println("\noutput:", Serialize(ans))
}
