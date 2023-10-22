// Created by Yuan at 2023/10/22 11:26
// leetgo: dev
// https://leetcode.cn/problems/next-permutation/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseNums(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp

		i++
		j--
	}
}

func nextPermutation(nums []int) {
	// 从后往前找到第一个比前面一个小的数 A
	// 再从后往前找到第一个比上面找到的那个数要小的第一个数 B
	// 将 A、B 互换位置所得的数即为比原数要大“一点”的数

	// 因为是原地修改的数组，一定要从低位找一个大数替换高位的小数
	// 要想有序的依次增大原数，高位小数和低位大数一定挨得最近

	// 假设 A 前面的数为 K
	// nums[-1] ~ nums[k] 为递增序列
	// nums[-1] ~ nums[Bi] 肯定也为递增
	// nums[Bi] 是第一个比 nums[Ai] 小的数
	//	所以 nums[-1] ~ nums[Ai] 也是递增，nums[Ai] ~ nums[k] 也是递增的
	//	翻转后对 nums[-1] ~ nums[Ai] 倒序
	//	=> 最后原数组的下一个序列

	lastIdx := len(nums) - 1
	var pre int

	// 找 A
	pre = nums[lastIdx]
	var aIdx int = lastIdx
	for ; aIdx >= 0; aIdx-- {
		if nums[aIdx] < pre {
			break
		}
		pre = nums[aIdx]
	}

	if aIdx == -1 {
		// 整个数组倒序
		reverseNums(nums)
		return
	}

	A := nums[aIdx]

	// 找 B
	var bIdx int = lastIdx
	for ; bIdx >= 0; bIdx-- {
		if nums[bIdx] > A {
			break
		}
		pre = nums[bIdx]
	}
	B := nums[bIdx]

	nums[aIdx] = B
	nums[bIdx] = A

	reverseNums(nums[aIdx+1:])
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	nextPermutation(nums)
	ans := nums

	fmt.Println("\noutput:", Serialize(ans))
}
