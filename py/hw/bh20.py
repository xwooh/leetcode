"""
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。



示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]
"""

from collections import deque
from typing import Deque


class Solution:
    def maxSlidingWindow(self, nums: list[int], k: int) -> list[int]:
        # 关键: 维护一个单调队列
        ans: list[int] = []

        # [(idx, v)]
        q: Deque[int] = deque()
        for i in range(k):
            q = self.addQ(q, 0, i, nums)
        ans.append(nums[q[0]])

        for i in range(1, len(nums) - k + 1):
            q = self.addQ(q, i, i + k - 1, nums)
            ans.append(nums[q[0]])

        return ans

    def addQ(self, q: Deque[int], head: int, i: int, nums: list[int]) -> Deque[int]:
        # 队列能的数据表现形式
        # 0. 保证元素都在滑窗内
        # 1. 如果要加进来的元素比队尾大，弹出队尾元素
        # 2. 保证队头元素最大

        while q and q[0] < head:
            q.popleft()

        while q and nums[i] >= nums[q[-1]]:
            # 保证队尾永远最大
            q.pop()

        q.append(i)

        while len(q) > 1 and nums[q[0]] <= nums[q[1]]:
            q.popleft()

        return q


if __name__ == "__main__":
    print(
        Solution().maxSlidingWindow([1, 3, -1, -3, 5, 3, 6, 7], 3) == [3, 3, 5, 5, 6, 7]
    )
    print(Solution().maxSlidingWindow([1, 3, 1, 2, 0, 5], 3) == [3, 3, 2, 5])
