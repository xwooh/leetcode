// Created by Yuan at 2024/02/05 14:34
// leetgo: dev
// https://leetcode.cn/problems/permutation-in-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func check(m1, m2 map[string]int) bool {
	for k, c := range m1 {
		if c > m2[k] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	sm1 := map[string]int{}
	for _, c1 := range s1 {
		sm1[string(c1)] += 1
	}

	sm2 := map[string]int{}
	for i := 0; i < len(s1); i++ {
		sm2[string(s2[i])] += 1
	}

	l, e := 0, len(s2)-len(s1)
	for l <= e {
		if l != 0 {
			// 目前的索引范围是 [l, l+len(s1)-1]
			// l-1 出去了，进来一个字符是 l+len(s1)
			sm2[string(s2[l-1])] -= 1
			sm2[string(s2[l+len(s1)-1])] += 1
		}

		if check(sm1, sm2) {
			return true
		}

		l++
	}
	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s1 := Deserialize[string](ReadLine(stdin))
	s2 := Deserialize[string](ReadLine(stdin))
	ans := checkInclusion(s1, s2)

	fmt.Println("\noutput:", Serialize(ans))
}
