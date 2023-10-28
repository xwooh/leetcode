// Created by Yuan at 2023/10/28 18:36
// leetgo: dev
// https://leetcode.cn/problems/minimum-add-to-make-parentheses-valid/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func minAddToMakeValid(s string) (ans int) {

	// rn 的需求数
	var rn int

	for _, c := range s {
		if string(c) == ")" {
			rn--
			if rn == -1 {
				// 说明 ) 超了，需要插入一个左括号
				ans++
				rn = 0
			}
		} else if string(c) == "(" {
			rn++
		}
	}

	ans += rn

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := minAddToMakeValid(s)

	fmt.Println("\noutput:", Serialize(ans))
}
