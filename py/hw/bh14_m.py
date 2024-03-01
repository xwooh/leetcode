"""
⭐️⭐️⭐️ 373. 查找和最小的 K 对数字

给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。


示例 1:
输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

示例 2:
输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
输出: [1,1],[1,1]
解释: 返回序列中的前 2 对数：
     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
"""

import heapq


class Solution:
    def kSmallestPairs(
        self, nums1: list[int], nums2: list[int], k: int
    ) -> list[list[int]]:
        """
        (0, 0) (0, 1) ...
        (1, 0) (1, 1) ...
        (2, 0) (1, 1) ...

        每行行头数组和肯定是这行中最小的值，先把每行行头入堆，取出一个后把该行后一个数加入

        因为这样能保证堆里面的数组和就是当前最小的几个数
        """
        ans = []

        line_head = [[nums1[i] + nums2[0], i, 0] for i in range(len(nums1))]
        heapq.heapify(line_head)

        while k > 0:
            x = heapq.heappop(line_head)
            i, j = x[1], x[2]

            ans.append([nums1[i], nums2[j]])

            if j < len(nums2) - 1:
                # 塞 i, j+1
                heapq.heappush(line_head, [nums1[i] + nums2[j + 1], i, j + 1])

            k -= 1

        return ans


if __name__ == "__main__":
    print(Solution().kSmallestPairs([1, 1, 2], [1, 2, 3], 2))
    print(Solution().kSmallestPairs([1, 7, 11], [2, 4, 6], 5))
    print(Solution().kSmallestPairs([1, 2, 4, 5, 6], [3, 5, 7, 9], 3))
