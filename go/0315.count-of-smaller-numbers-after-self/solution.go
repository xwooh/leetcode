// Created by Yuan at 2023/12/19 17:28
// leetgo: dev
// https://leetcode.cn/problems/count-of-smaller-numbers-after-self/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type Vc struct {
	idx int
	v   int
	c   int
}

func mergeTwo(ns1 []Vc, ns2 []Vc) []Vc {
	var p1, p2 int
	var ns []Vc

	len1 := len(ns1)
	len2 := len(ns2)

	for p1 < len1 && p2 < len2 {
		if ns1[p1].v > ns2[p2].v {
			ns1[p1].c += len2 - p2
			ns = append(ns, ns1[p1])
			p1++
		} else {
			ns = append(ns, ns2[p2])
			p2++
		}
	}

	ns = append(ns, ns2[p2:]...)
	ns = append(ns, ns1[p1:]...)

	return ns
}

func mergeDesc(nums []Vc) []Vc {
	if len(nums) <= 1 {
		return nums
	}

	lns := mergeDesc(nums[:len(nums)/2])
	rns := mergeDesc(nums[len(nums)/2:])

	ns := mergeTwo(lns, rns)
	return ns
}

func countSmaller(nums []int) (ans []int) {
	if len(nums) == 0 {
		return
	} else if len(nums) == 1 {
		ans = []int{0}
		return
	}

	// 初始化
	ans = make([]int, len(nums))
	vcs := []Vc{}
	for idx, v := range nums {
		vcs = append(vcs, Vc{idx, v, 0})
	}

	vcs = mergeDesc(vcs)
	for _, vc := range vcs {
		ans[vc.idx] = vc.c
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := countSmaller(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
