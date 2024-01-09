// Created by Yuan at 2024/01/09 11:37
// leetgo: dev
// https://leetcode.cn/problems/remove-k-digits/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeKdigits(num string, k int) string {
	// 单调栈，只要新元素比栈顶元素小，就弹出栈顶元素

	// 都会被弹掉，直接返回 0
	if len(num) <= k {
		return "0"
	}

	stack := []string{}
	for _, n := range num {
		s := string(n)

		// 判断是否比栈顶元素小
		for len(stack) > 0 && s < stack[len(stack)-1] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}

		stack = append(stack, s)
	}

	// 如果是非递减的数，那上述不会有任何出栈操作
	//	所以需要从后往前弹出
	for len(stack) > 0 && k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	// 都弹出去了，就返回 0
	if len(stack) == 0 {
		return "0"
	}

	// 把前面的 0 去掉
	s := strings.Join(stack, "")
	for strings.HasPrefix(s, "0") {
		s = strings.TrimPrefix(s, "0")
	}

	if s == "" {
		return "0"
	}

	// 还有值就直接返回
	return s
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	num := Deserialize[string](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := removeKdigits(num, k)

	fmt.Println("\noutput:", Serialize(ans))
}
