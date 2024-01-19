// Created by Yuan at 2024/01/19 13:26
// leetgo: dev
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeDuplicates(nums []int) (ans int) {
	// 当前元素计数超过两次后开始数后面有多少个重复的元素 (记为x)
	// 从下一个不重复元素开始往前挪 x 位

	dc := 0
	nc := 0
	e := math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] == e {
			// 当前元素和前一个元素相同
			if dc == 2 {
				// 前面已经有 2 个一样的元素了
				nc++
			} else {
				// 还不够两个相同元素
				if ans != i {
					nums[ans] = nums[i]
				}

				dc++
				ans++
			}
		} else {
			// 当前元素和前面元素不一样
			dc = 1 // 重复计数重置为 1
			nc = 0 // 多余元素计数重置为 0

			// 有重复多余的元素
			// 覆盖掉这些多余元素
			if ans != i {
				nums[ans] = nums[i]
			}

			ans++ // 有效元素++
		}
		e = nums[i]
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := removeDuplicates(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
