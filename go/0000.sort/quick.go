package main

func swap(nums []int, a, b int) {
	tmp := nums[b]
	nums[b] = nums[a]
	nums[a] = tmp
}

func partition(nums []int, start, end int) (pIdx int) {
	pIdx = start
	p := nums[pIdx]

	for pIdx <= end {
		if nums[pIdx] < p {
			swap(nums, pIdx, start)
			pIdx++  // 此时 pIdx 位置的值是由 start 位置换过来的，所以肯定 <= p，pIdx++ 去检查下一个值
			start++ // start 位置上换上了新值，++ 去拿 pIdx 位置这个坑位
		} else if nums[pIdx] > p {
			swap(nums, pIdx, end)
			end-- // end 位置上换上了新值，-- 去占下一个坑
			// pIdx 上换到的值还没判断过，pIdx 不能动
		} else {
			pIdx++
		}
	}
	pIdx--
	return
}

func quick(nums []int) {
	// 取一个数，不大于它的放左边，大于它的放右边
	// 然后递归它的左边和右边
	if len(nums) <= 1 {
		return
	}

	pIdx := partition(nums, 0, len(nums)-1)
	// fmt.Printf("%v, %d(%d), %v\n", nums[:pIdx], nums[pIdx], pIdx, nums[pIdx+1:])
	quick(nums[:pIdx])
	quick(nums[pIdx+1:])

	return
}

func topK(nums []int, k int, ans *[]int) {
	p := partition(nums, 0, len(nums)-1)

	if p+1 == k {
		*ans = append(*ans, nums[:p+1]...)
		return
	} else if p+1 > k {
		topK(nums[:p], k, ans)
		return
	} else if p+1 < k {
		*ans = append(*ans, nums[:p+1]...)
		topK(nums[p+1:], k-(p+1), ans)
		return
	}
}
