"""
求满足条件的最长子串的长度 滑动窗口

题目描述

给定一个字符串，只包含字母和数字，按要求找出字符串中的最长(连续)子串的长度，字符串本身是其最长的子串，子串要求:

1、只包含 1 个字母(a~z,A~Z)，其余必须是数字:

2、字母可以在子串中的任意位置;

如果找不到满足要求的子串，如全是字母或全是数字，则返回-1。

输入描述

字符串(只包含字母和数字)

输出描述

子串的长度
"""


def count(s: str) -> int:
    """
    r 往前走，一旦 wf、nf 都为 True 之后，停止往前
    l 前进道 r-1 处，重置重新计数
    """
    ans = 0

    l, r = 0, 0

    wc = 0
    nc = 0
    while r < len(s):
        if s[r].isdigit():
            nc += 1
        else:
            wc += 1

        if wc >= 2 and nc >= 2:
            # 要重置了
            wc = 0
            nc = 0
            ans = max(r - l, ans)
            l = r - 1
            r = l
        else:
            r += 1
            ans = max(r - l, ans)

    return ans


if __name__ == "__main__":
    print(count("abC124ACb"))  # 4
    print(count("a5"))  # 2
    print(count("a528372adksjsj9"))  # 8
