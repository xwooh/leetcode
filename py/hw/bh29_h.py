"""
书籍叠放 最长递增子序列+二分查找

题目描述

书籍的长、宽都是整数对应(l,w)。如果书 A 的长宽度都比 B 长宽大时，则允许将 B 排列

放在 A 上面。现在有一组规格的书籍，书籍叠放时要求书籍不能做旋转，请计算最多能有

多少个规格书籍能叠放在一起。


输入描述

输入: books = [[20,16],[15,11],[10,10],[9,10]]

说明: 总共 4 本书籍，第一本长度为 20 宽度为 16;第二本书长度为 15 宽度为 11，依次

类推，最后一本书长度为 9 宽度为 10.


输出描述

输出: 3
说明: 最多 3 个规格的书籍可以叠放到一起,从下到上依次为: [20,16],[15,11],[10,10]
"""


def lbs(nums: list[int], target: int) -> int:
    l, r = 0, len(nums) - 1

    while l < r:
        m = l + (r - l) // 2
        if nums[m] >= target:
            r = m
        else:
            l = m + 1

    return l


def count(books: list[list]) -> int:
    if len(books) <= 1:
        return len(books)

    books.sort(key=lambda x: (x[0], -x[1]))

    p = [books[0][1]]
    for book in books[1:]:
        if book[1] > p[-1]:
            p.append(book[1])
        else:
            # 因为 book[1] 比 p[-1] 小，所以必定能找到一个左边界
            idx = lbs(p, book[1])
            p[idx] = book[1]

    return len(p)


if __name__ == "__main__":
    print(count([[20, 16], [15, 11], [10, 10], [13, 11], [9, 10], [14, 12]]))
