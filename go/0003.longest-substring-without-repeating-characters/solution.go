// Created by Yuan at 2023/07/23 19:33
// leetgo: dev
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLongestSubstring(s string) (ans int) {
	seen := make(map[rune]int)

	var start int
	for end, w := range s {
		eIdx, in := seen[w]
		if in && eIdx >= start {
			// 如果存在，而且在当前 [start, end] 中，那就从这个字符后面开始重新算 unique 长度
			start = eIdx + 1
		}

		seen[w] = end

		cl := end - start + 1
		if cl > ans {
			ans = cl
		}
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLongestSubstring(s)

	fmt.Println("\noutput:", Serialize(ans))
}
