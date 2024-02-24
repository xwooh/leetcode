// Created by Yuan at 2024/02/24 19:52
// leetgo: dev
// https://leetcode.cn/problems/sum-of-square-numbers/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func judgeSquareSum(c int) bool {
	m := int(math.Sqrt(float64(c)))

	for a := 0; a <= m; a++ {
		fb := math.Sqrt(float64(c - a*a))
		if fb == math.Floor(fb) {
			return true
		}
	}

	return false

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	c := Deserialize[int](ReadLine(stdin))
	ans := judgeSquareSum(c)

	fmt.Println("\noutput:", Serialize(ans))
}
