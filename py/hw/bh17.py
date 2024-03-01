"""
非严格递增连续数字序列

题目描述：
输入一个字符串仅包含大小写字母和数字，
求字符串中包含的最长的非严格递增连续数字序列的长度（比如12234属于非严格递增连续数字序列）。

输入描述:
输入一个字符串仅包含大小写字母和数字，输入的字符串最大不超过255个字符。

输出描述：
最长的非严格递增连续数字序列的长度
"""


def count(s: str) -> int:
    ml = 0

    pre: None | int = None
    cl = 0
    for c in s:
        if not c.isdigit():
            continue
        cn = int(c)
        if pre is None:
            pre = cn
            cl += 1
            ml = max(ml, cl)
        elif cn >= pre:
            pre = cn
            cl += 1
            ml = max(ml, cl)
        else:
            # 以当前数字为起始值
            pre = cn
            cl = 1

    return ml


if __name__ == "__main__":
    print(count("abc2234019A334bc"))  # 4
    print(count("aaaaaa44ko543j123j7345677781"))  # 8
    print(count("345678a44ko543j123j7134567778aa"))  # 9
