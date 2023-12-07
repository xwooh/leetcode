// Created by Yuan at 2023/11/23 10:54
// leetgo: dev
// https://leetcode.cn/problems/longest-common-prefix/

package main

import (
	"bufio"
	"fmt"
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

func calcLongestCommonPrefix(strA, strB string) int {
	var i int
	for i < len(strA) && i < len(strB) {
		if strA[i] == strB[i] {
			i++
		} else {
			break
		}
	}
	return i - 1
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	pre := strs[0]
	i := 1
	for ; i < len(strs); i++ {
		idx := calcLongestCommonPrefix(pre, strs[i])
		if idx < 0 {
			return ""
		}
		pre = strs[i][:idx+1]
	}

	return pre
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	strs := Deserialize[[]string](ReadLine(stdin))
	ans := longestCommonPrefix(strs)

	fmt.Println("\noutput:", Serialize(ans))
}
