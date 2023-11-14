// Created by Yuan at 2023/09/01 19:09
// leetgo: dev
// https://leetcode.cn/problems/3sum/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func swap(nums []int, a, b int) {
	tmp := nums[b]
	nums[b] = nums[a]
	nums[a] = tmp
}

func partition(nums []int, start, end int) int {
	pIdx := start
	p := nums[pIdx]

	for pIdx <= end {
		if nums[pIdx] > p {
			swap(nums, pIdx, end)
			end--
		} else if nums[pIdx] < p {
			swap(nums, pIdx, start)
			start++
			pIdx++
		} else {
			pIdx++
		}
	}

	pIdx--

	return pIdx
}

func qs(nums []int) {
	if len(nums) <= 1 {
		return
	}

	p := partition(nums, 0, len(nums)-1)
	qs(nums[:p])
	qs(nums[p+1:])

	return
}

func bs(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		m := nums[mid]
		if m == target {
			return mid
		} else if m > target {
			r = mid - 1
		} else if m < target {
			l = mid + 1
		}
	}

	return -1
}

func threeSum(nums []int) (ans [][]int) {
	// 先排序
	// 双指针指向队头队尾
	// 对双指针中间元素通过二分找 target

	if len(nums) < 3 {
		return [][]int{}
	}

	// 排序
	qs(nums)

	// 这边不用 map：会超时过不了
	preI := math.MaxInt
	// 双指针二分找 target
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > 0 {
			// 最小的都 > 0 了，那不可能三数之和等于 0
			break
		}

		if nums[i] == preI {
			continue
		}
		preI = nums[i]

		preJ := math.MaxInt
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < 0 {
				// 最大值都 < 0 了，那不可能三数之和等于 0
				break
			}

			if nums[j] == preJ {
				continue
			}
			preJ = nums[j]

			start := i + 1
			t := bs(nums[start:j], 0-nums[i]-nums[j])
			if t != -1 {
				ans = append(ans, []int{nums[i], nums[start+t], nums[j]})
			}
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := threeSum(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
