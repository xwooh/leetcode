package main

func mergeTwo(nsa []int, nsb []int) (ans []int) {
	if len(nsa) == 0 {
		ans = nsb
		return
	}
	if len(nsb) == 0 {
		ans = nsa
		return
	}

	var i, j int
	for i < len(nsa) && j < len(nsb) {
		if nsa[i] <= nsb[j] {
			ans = append(ans, nsa[i])
			i++
		} else {
			ans = append(ans, nsb[j])
			j++
		}
	}

	if i == len(nsa) {
		ans = append(ans, nsb[j:]...)
	} else {
		ans = append(ans, nsa[i:]...)
	}

	return
}

func merge(nums []int) (ans []int) {
	if len(nums) <= 1 {
		return nums
	} else if len(nums) == 2 {
		if nums[0] <= nums[1] {
			return nums
		} else {
			return []int{nums[1], nums[0]}
		}
	}

	lns := merge(nums[:len(nums)/2])
	rns := merge(nums[len(nums)/2:])

	ans = mergeTwo(lns, rns)

	return
}
