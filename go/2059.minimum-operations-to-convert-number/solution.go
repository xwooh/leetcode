// Created by Yuan at 2024/01/25 11:54
// leetgo: dev
// https://leetcode.cn/problems/minimum-operations-to-convert-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minimumOperations(nums []int, start int, goal int) (ans int) {
	// BFS 解法

	solutions := []int{start}
	// 已经出现过的结果
	views := map[int]int{}
	for len(solutions) > 0 {

		// 题目指明 start != goal，故省略这个判断
		// if slices.Contains(solutions, goal) {
		// 	return
		// }

		ans++
		// 对于 solutions 中的所有元素进行一遍 +/-/^ 操作，看看结果里面有没有 goal
		// 一次这种操作即为一次操作数
		ns := []int{}
		for _, s := range solutions {
			for _, v := range nums {
				sp := s + v
				sm := s - v
				se := s ^ v

				// x 在 [0, 1000] 之间才能再做下一次运算操作
				// && 之前没出现过的数才能进行下一次运算
				if sp >= 0 && sp <= 1000 && views[sp] == 0 {
					ns = append(ns, sp)
				}
				views[sp] = 1

				if sm >= 0 && sm <= 1000 && views[sm] == 0 {
					ns = append(ns, sm)
				}
				views[sm] = 1

				if se >= 0 && se <= 1000 && views[se] == 0 {
					ns = append(ns, se)
				}
				views[se] = 1
			}
		}

		if views[goal] == 1 {
			return
		}

		solutions = ns

	}

	return -1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	start := Deserialize[int](ReadLine(stdin))
	goal := Deserialize[int](ReadLine(stdin))
	ans := minimumOperations(nums, start, goal)

	fmt.Println("\noutput:", Serialize(ans))
}
