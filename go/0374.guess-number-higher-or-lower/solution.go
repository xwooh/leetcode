// Created by Yuan at 2023/10/27 17:45
// leetgo: dev
// https://leetcode.cn/problems/guess-number-higher-or-lower/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

func guess(num int) int

// @lc code=begin

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) (ans int) {

	var l, r = 1, n
	for l <= r {
		ans = (l + r) / 2
		j := guess(ans)
		if j == 0 {
			return
		} else if j > 0 {
			l = ans + 1
		} else if j < 0 {
			r = ans - 1
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	pick := Deserialize[int](ReadLine(stdin))
	ans := guessNumber(n, pick)

	fmt.Println("\noutput:", Serialize(ans))
}
