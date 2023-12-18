package main

func bsl(nums []int, target int) int {
	// 二分搜索左边界
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2
		if nums[m] >= target {
			// 往左缩
			r = m
		} else if nums[m] < target {
			// 往右缩
			l = m + 1
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}

func bsll(nums []int, target int) int {
	// 找比 target 小的第一个元素的索引
	l, r := 0, len(nums)-1

	for l < r {
		// NOTE: 往右边缩的时候没有 m+1，所以 m 要往右偏一点
		m := l + (r-l)/2 + 1
		if nums[m] >= target {
			// 当前元素大，那就往左边找
			r = m - 1
		} else {
			l = m
		}
	}

	if nums[l] < target {
		return l
	}

	return -1
}

func bsr(nums []int, target int) int {
	// 二分搜索右边界
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2 + 1
		if nums[m] == target {
			// 找右边界，左移
			l = m
		} else if nums[m] > target {
			// 左移
			r = m - 1
		} else if nums[m] < target {
			// 右移
			l = m + 1
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}

func bsrr(nums []int, target int) int {
	// 找比 target 大的第一个元素的索引
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2
		if nums[m] <= target {
			// 当前小，那就右移
			l = m + 1
		} else {
			r = m
		}
	}

	if nums[l] > target {
		return l
	}

	return -1
}
