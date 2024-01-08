// Created by Yuan at 2023/11/15 13:44
// leetgo: dev
// https://leetcode.cn/problems/top-k-frequent-words/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

func swap(wcs []WC, a, b int) {
	tmp := wcs[a]
	wcs[a] = wcs[b]
	wcs[b] = tmp
}

func partition(wcs []WC, start, end int) int {
	pIdx := start
	p := wcs[pIdx]

	for pIdx <= end {
		if wcs[pIdx].c > p.c {
			swap(wcs, pIdx, start)
			start++
			pIdx++
		} else if wcs[pIdx].c < p.c {
			swap(wcs, pIdx, end)
			end--
		} else {
			pIdx++
		}
	}

	pIdx--
	return pIdx
}

func topK(wcs []WC, k int, ans *[]WC) {
	p := partition(wcs, 0, len(wcs)-1)

	if p+1 == k {
		*ans = append(*ans, wcs[:p+1]...)
		return
	} else if p+1 > k {
		topK(wcs[:p], k, ans)
		return
	} else if p+1 < k {
		*ans = append(*ans, wcs[:p+1]...)
		topK(wcs[p+1:], k-(p+1), ans)
	}
}

// @lc code=begin
type WC struct {
	w string
	c int
}

func isj(i, j string) bool {
	// 字符串 i 是不是比字符串 j 的字典顺序小
	var wi, wj int
	for wi < len(i) && wj < len(j) {
		if i[wi] < j[wj] {
			return true
		} else if i[wi] > j[wj] {
			return false
		} else {
			wi++
			wj++
		}
	}
	if wi == len(i) {
		// j 的单词长，那就是 i 小
		return true
	}
	return false
}

func msort(lwcs []WC, rwcs []WC) (ans []WC) {
	if len(lwcs) == 0 {
		return rwcs
	}
	if len(rwcs) == 0 {
		return rwcs
	}

	var i, j int
	for i < len(lwcs) && j < len(rwcs) {
		if lwcs[i].c < rwcs[j].c {
			ans = append(ans, rwcs[j])
			j++
		} else if lwcs[i].c > rwcs[j].c {
			ans = append(ans, lwcs[i])
			i++
		} else if lwcs[i].c == rwcs[j].c {
			// 按单词的字典序排序
			if isj(lwcs[i].w, rwcs[j].w) {
				ans = append(ans, lwcs[i])
				i++
			} else {
				ans = append(ans, rwcs[j])
				j++
			}
		}
	}

	if i == len(lwcs) {
		ans = append(ans, rwcs[j:]...)
	} else if j == len(rwcs) {
		ans = append(ans, lwcs[i:]...)
	}

	return
}

func sort(wcs []WC) []WC {
	if len(wcs) <= 1 {
		return wcs
	} else if len(wcs) == 2 {
		if wcs[0].c > wcs[1].c {
			return wcs
		} else if wcs[0].c < wcs[1].c {
			return []WC{wcs[1], wcs[0]}
		} else {
			if isj(wcs[0].w, wcs[1].w) {
				return wcs
			} else {
				return []WC{wcs[1], wcs[0]}
			}
		}
	}

	lwcs := sort(wcs[:len(wcs)/2])
	rwcs := sort(wcs[len(wcs)/2:])

	return msort(lwcs, rwcs)
}

func topKFrequent(words []string, k int) (ans []string) {

	// 保证按照 word 出现顺序记录
	wl := []WC{}

	// 统计频率
	m := map[string]int{}
	for _, s := range words {
		if _, ok := m[s]; ok {
			wl[m[s]].c++
		} else {
			// 此时 wl 的末尾位置
			m[s] = len(wl)
			wl = append(wl, WC{s, 1})
		}
	}

	// 按频率排序
	wl = sort(wl)
	for _, t := range wl {
		if len(ans) < k {
			ans = append(ans, t.w)
		} else {
			break
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	words := Deserialize[[]string](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := topKFrequent(words, k)

	fmt.Println("\noutput:", Serialize(ans))
}
