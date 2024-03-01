"""
组成最大数

题目描述

小组中每位都有一张卡片，卡片上是6位内的正整数，将卡片连起来可以组成多种数字，计算组成的最大数字。

输入描述

“，”号分割的多个正整数字符串，不需要考虑非数字异常情况，小组最多25个人。

输出描述

最大的数字字符串
"""


def djuge(l: str, r: str) -> int:
    """
    -1: l
     0: =
     1: r
    """
    idx = 0
    l_len = len(l)
    r_len = len(r)
    end = max(l_len, r_len)

    while idx < end:

        # 如果位数不够，就用最后一个数补全
        if idx >= l_len:
            i_n = int(l[-1])
        else:
            i_n = int(l[idx])

        if idx >= r_len:
            j_n = int(r[-1])
        else:
            j_n = int(r[idx])

        if i_n > j_n:
            return -1
        elif j_n > i_n:
            return 1
        else:
            idx += 1

    return 0


def djuge2(l: str, r: str) -> int:
    """
    -1: l
     0: =
     1: r
    """
    lr = int(l) * 10 ** len(r) + int(r)
    rl = int(r) * 10 ** len(l) + int(l)

    if lr > rl:
        return -1
    elif rl > lr:
        return 1
    return 0


def merge(ls: list[str], rs: list[str]) -> list[str]:
    ns = []

    i, j = 0, 0

    while i < len(ls) and j < len(rs):
        r = djuge2(ls[i], rs[j])
        if r == -1:
            ns.append(ls[i])
            i += 1
        elif r == 1:
            ns.append(rs[j])
            j += 1
        else:
            ns.append(ls[i])
            ns.append(rs[j])
            i += 1
            j += 1

    if i == len(ls):
        ns += rs[j:]
    else:
        ns += ls[i:]

    return ns


def sort_str(ns: list[str]) -> list[str]:
    if len(ns) <= 1:
        return ns

    m = len(ns) // 2

    l = sort_str(ns[:m])
    r = sort_str(ns[m:])

    return merge(l, r)


def transfer(s: str) -> str:
    """
    10,2 -> 210

    对所有的字符串排序，排序规则：
    数字越大排的越前面
    """

    ns = sort_str(s.split(","))
    r = "".join(ns)
    if r.startswith("0"):
        return "0"
    return "".join(ns)


if __name__ == "__main__":
    print(transfer("10,2"))
    print(transfer("34323,3432"))
    print(transfer("10,20,11,12"))
    print(transfer("3,30,34,5,9"))
    print(transfer("22,221,33,93,96,933,98,973,974,9,9999"))
    print(transfer("98,973,974,9,9999"))
