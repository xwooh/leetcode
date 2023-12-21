// Created by Yuan at 2023/12/20 18:08
// leetgo: dev
// https://leetcode.cn/problems/merge-intervals/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type Interv [][]int

func (intv Interv) Len() int {
	return len(intv)
}

func (intv Interv) Less(i, j int) bool {
	return intv[i][0] < intv[j][0]
}

func (intv Interv) Swap(i, j int) {
	intv[i], intv[j] = intv[j], intv[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) (ans [][]int) {
	sort.Sort(Interv(intervals))

	for _, ns := range intervals {
		if len(ans) == 0 || ns[0] > ans[len(ans)-1][1] {
			ans = append(ans, ns)
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], ns[1])
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	intervals := Deserialize[[][]int](ReadLine(stdin))
	ans := merge(intervals)

	fmt.Println("\noutput:", Serialize(ans))
}
