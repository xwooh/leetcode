// Created by Yuan at 2024/02/19 11:49
// leetgo: dev
// https://leetcode.cn/problems/subsets/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func dfs(startIdx int, nums []int, path []int, k int, ans *[][]int) {
	if len(path) == k {
		np := make([]int, len(path))
		copy(np, path)
		*ans = append(*ans, np)
		return
	}

	for i := startIdx; len(nums)-i >= k-len(path); i++ {
		path = append(path, nums[i])
		dfs(i+1, nums, path, k, ans)
		path = path[:len(path)-1]
	}
}

func pick(nums []int, k int, ans *[][]int) {
	dfs(0, nums, []int{}, k, ans)
	return
}

func subsets(nums []int) (ans [][]int) {
	ans = append(ans, []int{})

	for i := 1; i <= len(nums); i++ {
		pick(nums, i, &ans)
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := subsets(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
