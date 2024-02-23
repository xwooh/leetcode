// Created by Yuan at 2024/02/23 14:28
// leetgo: dev
// https://leetcode.cn/problems/rotate-image/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func rotateOne(matrix [][]int, i, j int) {
	// 从 (i, j) 开始挨个旋转移动

	delta := len(matrix) - 1

	lastV := math.MaxInt
	// 每次旋转都要变动 4 个点（有4条边）
	for t := 4; t > 0; t-- {
		// 转换一下坐标位置
		x, y := i, j+delta
		// 旋转后 (x, y) 的位置
		targetX, targetY := y, -x
		// 该点实际在元素组中的位置
		inX, inY := targetX, targetY+delta
		// 再转换到实际数组中的索引位置
		mI, mJ := inX-delta, inY

		// fmt.Printf("i: %d, j: %d, mX: %d, mY: %d\n", i, j, mI, mJ)

		if lastV == math.MaxInt {
			lastV = matrix[i][j]
		}

		if lastV == math.MaxInt {
			lastV = matrix[i][j]
		}
		// 保存被替换到那个位置的元素
		tmp := matrix[mI][mJ]
		// 旋转当前位置的元素到指定位置上
		matrix[mI][mJ] = lastV

		lastV = tmp
		// 对替换的位置进行下一轮替换
		i, j = mI, mJ
	}
}

func rotate(matrix [][]int) {
	n := len(matrix)
	b := n - 1
	if n == 2 {
		// 二维数组，至少要转一次
		b = 1
	} else if n%2 == 0 {
		// 偶数，少转一层
		b = b - 1
	}

	for i := 0; i < b; i++ {
		for j := i; j < n-1-i; j++ {
			rotateOne(matrix, i, j)
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	rotate(matrix)
	ans := matrix

	fmt.Println("\noutput:", Serialize(ans))
}
