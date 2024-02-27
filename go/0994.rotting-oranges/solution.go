// Created by Yuan at 2024/02/27 08:58
// leetgo: dev
// https://leetcode.cn/problems/rotting-oranges/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type P struct {
	x, y int
}

func orangesRotting(grid [][]int) (ans int) {
	m := len(grid)
	n := len(grid[0])

	s := []P{}
	freshC := 0
	// 统计一共多少新鲜橘子，腐烂的橘子入队
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				freshC++
			} else if grid[i][j] == 2 {
				s = append(s, P{i, j})
			}
		}
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for {
		for c := len(s); c > 0; c-- {
			p := s[0]
			s = s[1:]

			for _, d := range directions {
				ni, nj := p.x+d[0], p.y+d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 1 {
					// 感染新鲜的
					grid[ni][nj] = 2
					freshC--

					// 去感染下一轮
					s = append(s, P{ni, nj})
				}
			}
		}

		if len(s) <= 0 {
			break
		}

		ans++
	}

	// 如果还有新鲜的，那就是没法全部腐烂
	if freshC > 0 {
		return -1
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	grid := Deserialize[[][]int](ReadLine(stdin))
	ans := orangesRotting(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
