// Created by Yuan at 2023/10/17 22:51
// leetgo: dev
// https://leetcode.cn/problems/permutations/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func travel(selection []int, selected []int, used map[int]bool, res *[][]int) {
	if len(selected) == len(selection) {
		// 都选完了
		tmp := make([]int, len(selected))
		copy(tmp, selected)
		*res = append(*res, tmp)
		return
	}

	for _, v := range selection {
		if used[v] {
			continue
		}

		selected = append(selected, v)
		used[v] = true
		travel(selection, selected, used, res)
		selected = selected[:len(selected)-1]
		used[v] = false
	}
}

func permute(nums []int) (ans [][]int) {

	used := make(map[int]bool, len(nums))

	selected := make([]int, 0)
	// selected := []int{}

	travel(nums, selected, used, &ans)

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := permute(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
