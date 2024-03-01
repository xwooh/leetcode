class Solution:
    def lengthOfLISOn2(self, nums: list[int]) -> int:
        """
        定义 m[i] 为 i 结尾的最长递增子序列长度
        那这样的话 m[i+1] = max(
                                回溯找比 nums[i+1] 小的数结尾的最长连续子串的最大值
                            )
        """

        m = [1 for _ in range(len(nums))]

        for i in range(len(nums)):
            if i == 0:
                m[i] = 1
            else:
                # 找比 nums[i] 小的最大值
                j = i - 1
                mj = 0
                while j >= 0:
                    # 比当前数小的结尾的最大值
                    if nums[j] < nums[i]:
                        mj = max(mj, m[j])
                    j -= 1

                m[i] = mj + 1

        return max(m)

    def bs(self, nums: list[int], target: int) -> int:
        """二分找 >= target 的左界"""
        l, r = 0, len(nums) - 1
        while l < r:
            m = l + (r - l) // 2
            if nums[m] >= target:
                r = m
            else:
                l = m + 1
        return l

    def lengthOfLIS(self, nums: list[int]) -> int:
        """
        维护一个单调递增列表，后续 nums 中有元素插入到 >= x 的左界
        列表的长度就是上升子序列长度
        """
        p = [nums[0]]

        for n in nums[1:]:
            if n > p[-1]:
                p.append(n)
            else:
                idx = self.bs(p, n)
                p[idx] = n

        print(p)
        return len(p)


if __name__ == "__main__":
    print(Solution().lengthOfLIS([1, 3, 6, 7, 9, 4, 10, 5, 6]))
