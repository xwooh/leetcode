// Created by Yuan at 2023/10/26 17:06
// leetgo: dev
// https://leetcode.cn/problems/minimum-insertions-to-balance-a-parentheses-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minInsertions(s string) (ans int) {
	stack := []string{}

	for _, c := range s {
		if string(c) == "(" {
			stack = append(stack, "(")
		} else if string(c) == ")" {
			if len(stack) >= 2 && stack[len(stack)-1] == ")" && stack[len(stack)-2] == "(" {
				// 用当前的 ) 和栈最后的 ) 抵消栈倒数第二个 (
				stack = stack[:len(stack)-2]
				// 并且因为是右括号，插一个空格，防止和前面混在一起
				stack = append(stack, "-")
			} else {
				// 抵消不了就入栈
				stack = append(stack, ")")
			}
		}
		fmt.Printf("stack: %v\n", stack)
	}

	nr := false
	l, r := 0, len(stack)-1
	for l <= r {
		if l == r {
			if stack[l] == "(" || stack[r] == ")" {
				ans += 2
			}
			break
		}
		if stack[l] == ")" {
			if stack[r] == ")" {
				ans += 1
				l++
				r--
			} else if stack[r] == "(" {
				ans += 2
				r--
			} else if stack[r] == "-" {
				r--
			}
		} else if stack[l] == "(" {
			if stack[r] == ")" {
				if nr {
					r--
					nr = false
				} else {
					l++
					r--
					nr = true
					ans += 1
				}
			} else if stack[r] == "(" {
				ans += 2
				r--
			} else if stack[r] == "-" {
				r--
			}
		} else if stack[l] == "-" {
			l++
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := minInsertions(s)

	fmt.Println("\noutput:", Serialize(ans))
}
