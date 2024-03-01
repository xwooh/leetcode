// Created by Yuan at 2024/03/01 20:26
// leetgo: dev
// https://leetcode.cn/problems/russian-doll-envelopes/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func lbs(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		} else {
			return false
		}
	})

	p := []int{envelopes[0][1]}
	for _, e := range envelopes[1:] {
		if e[1] > p[len(p)-1] {
			p = append(p, e[1])
		} else {
			idx := lbs(p, e[1])
			p[idx] = e[1]
		}
	}

	return len(p)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	envelopes := Deserialize[[][]int](ReadLine(stdin))
	ans := maxEnvelopes(envelopes)

	fmt.Println("\noutput:", Serialize(ans))
}
