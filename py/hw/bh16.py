"""
乱序整数序列两数之和绝对值最小

题目描述

给定一个随机的整数(可能存在正整数和负整数)数组 nums，请你在该数组中找出两个

数，其和的绝对值(|nums[x]+nums[y]|)为最小值，并返回这个两个数 (按从小到大返回)以及

绝对值。

每种输入只会对应一个答案

但是，数组中同一个元素不能使用两遍

输入描述

个通过空格分割的有序整数席列字符串，最多 1000 个整数，目整数数值范围是

[-65535,65535].

输出描述

两数之和绝对值最小值
"""


def bs(nums: list[int], target: int) -> int:
    # 找 >= target 的左界
    l, r = 0, len(nums) - 1

    while l < r:
        m = l + (r - l) // 2
        if nums[m] < target:
            l = m + 1
        else:
            r = m

    return l


def find(nums: list[int]) -> list[int]:
    if len(nums) < 2:
        return []

    nums.sort()

    if nums[0] >= 0:
        # 全是正的
        return [nums[0], nums[1], nums[0] + nums[1]]
    elif nums[-1] <= 0:
        # 全是负的
        return [nums[-2], nums[-1], nums[-2] + nums[1]]

    # 有正有负，找最接近的
    ms = nums[-1] * 2
    n_idx, p_idx = 0, len(nums) - 1
    z_idx = bs(nums, 0)

    for ng_idx, ng in enumerate(nums[:z_idx]):
        # 找：这些负数和哪些整数之和绝对值最小
        idx = bs(nums, -ng)

        if ng_idx != idx:
            s = abs(ng + nums[idx])
            if s < ms:
                ms = s
                n_idx = ng_idx
                p_idx = idx

        if ng_idx != idx - 1:
            s = abs(ng + nums[idx - 1])
            if s < ms:
                ms = s
                n_idx = ng_idx
                p_idx = idx - 1

    return [nums[n_idx], nums[p_idx], ms]


if __name__ == "__main__":
    print(find([-1, -3, 7, 5, 11, 15]))
    print(find([-1, -3, 7, 5, 3, 15]))
    print(find([1, 3, 7, 5, 3, 15]))
    print(find([-1, -3, -7, -5, -3, 15]))
