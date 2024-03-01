"""
连续字母长度

题目描述

给定一个字符串，只包含大写字母，求在包含同一字母的子串中，长度第k长的子串的长度，相同字母只取最长的那个子串。

输入描述

第一行有一个子串(1<长度<=100)，只包含大写字母

第二行为k的值

输出描述

输出连续出现次数第k多的字母的次数

举例：
输入：AAAAHHHBBCDHHHH 3
输出：2
说明：
    同一字母连续出现的最多的是A和H，四次
    第二多的是H，3次，但是 H 已经存在 4 个连续的，顾不考虑
    下一个最长子串是 BB，所以答案输出2
"""


def count(s: str) -> dict[str, int]:
    # 关键：每个字母的最长长度
    m: dict[str, int] = {}
    pre = 0
    i = 1
    while i < len(s):
        if s[i] == s[pre]:
            i += 1
        else:
            if s[pre] in m:
                m[s[pre]] = max(m[s[pre]], i - pre)
            else:
                m[s[pre]] = i - pre
            pre = i
    return m


def find(s: str, k: int) -> int:
    m = count(s)
    ns = sorted(m.values())

    if len(ns) < k:
        return -1

    return ns[-k]


if __name__ == "__main__":
    print(find("AAAAHHHBBCDHHHH", 3))  # 2
    print(find("AABAAA", 2))  # 1
