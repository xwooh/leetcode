// Created by Yuan at 2023/10/23 21:16
// leetgo: dev
// https://leetcode.cn/problems/non-overlapping-intervals/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
import "sort"

func eraseOverlapIntervals(intervals [][]int) (ans int) {
	if len(intervals) <= 1 {
		return
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	lastB := intervals[0][0]
	lastE := intervals[0][1]
	for _, e := range intervals[1:] {
		if e[0] < lastE || e[0] < lastB || e[1] < lastE {
			ans++
		} else {
			lastB = e[0]
			lastE = e[1]
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	intervals := Deserialize[[][]int](ReadLine(stdin))
	ans := eraseOverlapIntervals(intervals)

	fmt.Println("\noutput:", Serialize(ans))
}
