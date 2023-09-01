// Created by Yuan at 2023/09/01 18:15
// leetgo: dev
// https://leetcode.cn/problems/number-of-ways-to-buy-pens-and-pencils/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func waysToBuyPensPencils(total int, cost1 int, cost2 int) (ans int64) {
	for i := 0; i <= total/cost1; i++ {
		// 计算剩余的钱能买几只 cost2 的笔
		// 0 只 cost2 的笔也算一种方案，1 只 cost2 的笔算 2 种方案，...
		ans += int64((total-(cost1*i))/cost2 + 1)
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	total := Deserialize[int](ReadLine(stdin))
	cost1 := Deserialize[int](ReadLine(stdin))
	cost2 := Deserialize[int](ReadLine(stdin))
	ans := waysToBuyPensPencils(total, cost1, cost2)

	fmt.Println("\noutput:", Serialize(ans))
}
