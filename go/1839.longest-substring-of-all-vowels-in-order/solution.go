// Created by Yuan at 2024/02/05 11:51
// leetgo: dev
// https://leetcode.cn/problems/longest-substring-of-all-vowels-in-order/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func toInt(s string) int {
	switch s {
	case "a":
		return 0
	case "e":
		return 1
	case "i":
		return 2
	case "o":
		return 3
	case "u":
		return 4
	default:
		return -1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestBeautifulSubstring(word string) (ans int) {
	st := []int{}

	for _, c := range word {
		i := toInt(string(c))
		if len(st) == 0 {
			if i == 0 {
				st = append(st, i)
			}
		} else if i-st[len(st)-1] == 0 || i-st[len(st)-1] == 1 {
			st = append(st, i)
		} else {
			// 当前的字母比 st 中结尾的字母要小，不符合要求
			// 因为题目要求所有字母至少出现一次，那就重新累计
			if i == 0 {
				st = []int{i}
			} else {
				st = []int{}
			}
		}

		var l int
		if len(st) < 5 {
			l = 0
		} else {
			if st[len(st)-1] != 4 {
				l = 0
			} else {
				l = len(st)
			}
		}
		ans = max(ans, l)
	}

	if ans < 5 {
		ans = 0
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	word := Deserialize[string](ReadLine(stdin))
	ans := longestBeautifulSubstring(word)

	fmt.Println("\noutput:", Serialize(ans))
}
