// Created by Yuan at 2023/11/07 21:45
// leetgo: dev
// https://leetcode.cn/problems/unique-paths/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func uniquePaths(m int, n int) (ans int) {

	p := make([][]int, m)
	for c := range p {
		p[c] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		p[0][i] = 1
	}
	for i := 0; i < m; i++ {
		p[i][0] = 1
	}

	for j := 1; j < m; j++ {
		for i := 1; i < n; i++ {
			p[j][i] = p[j-1][i] + p[j][i-1]
		}
	}

	ans = p[m-1][n-1]
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	m := Deserialize[int](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := uniquePaths(m, n)

	fmt.Println("\noutput:", Serialize(ans))
}
