// Created by Yuan at 2023/10/28 18:11
// leetgo: dev
// https://leetcode.cn/problems/missing-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func missingNumber(nums []int) (ans int) {

	s := 0
	for _, v := range nums {
		s += v
	}

	max := len(nums)
	ans = (1+max)*max/2 - s

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := missingNumber(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
