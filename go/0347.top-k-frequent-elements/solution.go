// Created by Yuan at 2023/11/14 13:35
// leetgo: dev
// https://leetcode.cn/problems/top-k-frequent-elements/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type E struct {
	v int
	c int
}

func swap(es []E, a, b int) {
	tmp := es[b]
	es[b] = es[a]
	es[a] = tmp
}

func partition(es []E, start, end int) int {
	// 大的放左边，小的放右边
	pIdx := start
	p := es[pIdx]

	for pIdx <= end {
		if es[pIdx].c > p.c {
			// 当前值比基准值大，放左边
			swap(es, pIdx, start)
			start++
			pIdx++
		} else if es[pIdx].c < p.c {
			// 当前值比基准值小，放右边
			swap(es, pIdx, end)
			end--
		} else {
			pIdx++
		}
	}
	pIdx--
	return pIdx
}

func topK(es []E, k int, ans *[]E) {
	p := partition(es, 0, len(es)-1)

	if p+1 == k {
		*ans = append(*ans, es[:p+1]...)
		return
	} else if p+1 > k {
		topK(es[:p], k, ans)
		return
	} else if p+1 < k {
		*ans = append(*ans, es[:p+1]...)
		topK(es[p+1:], k-(p+1), ans)
		return
	}
}

func topKFrequent(nums []int, k int) (ans []int) {
	m := map[int]int{}
	el := []E{}

	// 统计每个数的频率
	for _, v := range nums {
		if c, ok := m[v]; ok {
			m[v] = c + 1
		} else {
			m[v] = 1
		}
	}

	// 求频率最高的 topK
	for k, v := range m {
		el = append(el, E{k, v})
	}
	anEs := []E{}
	topK(el, k, &anEs)

	for _, e := range anEs {
		ans = append(ans, e.v)
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := topKFrequent(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
