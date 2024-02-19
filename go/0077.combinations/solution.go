// Created by Yuan at 2024/02/19 12:21
// leetgo: dev
// https://leetcode.cn/problems/combinations/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func pick(startIdx int, path []int, res *[][]int, n, k int) {
	// k, 取几个
	if len(path) == k {
		newPath := make([]int, len(path))
		copy(newPath, path)
		*res = append(*res, newPath)
		return
	}

	// 从 startIdx 开始，当剩余元素个数不足以满足 path = k 个时结束
	for i := startIdx; n-i+1 >= k-len(path); i++ {
		path = append(path, i)
		pick(i+1, path, res, n, k)
		path = path[:len(path)-1]
	}
}

func combine(n int, k int) (ans [][]int) {
	pick(1, []int{}, &ans, n, k)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := combine(n, k)

	fmt.Println("\noutput:", Serialize(ans))
}
