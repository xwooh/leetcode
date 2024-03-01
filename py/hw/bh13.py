"""
字符串加密

题目描述

给你一串末加密的字符串str，通过对字符串的每一个字母进行改变来实现加密，加密方式是在每一个字母str[i]偏移特定数组元素a[i]的量，数组a前三位已经赋值:

a[0]=1,a[1]=2,a[2]=4。

当i>=3时，数组元素a[i]=a[i-1]+al[i-2]+a[i-3]。

例如:原文abcde加密后bdgkr，其中偏移量分别是1,2.4,7,13。

输入描述

第一行为一个整数n(1<=n<=1000)，表示有n组测试数据，每组数据包含一行，原文str(只含有小写字母0<长度<=50)。

输出描述

每组测试数据输出一行，表示字符串的密文
"""

from typing import Iterator


def fib(k: int) -> Iterator[int]:
    f0, f1, f2 = 1, 2, 4

    ans = 0
    for idx in range(0, k + 1):
        if idx == 0:
            yield f0
        elif idx == 1:
            yield f1
        elif idx == 2:
            yield f2
        else:
            ans = f0 + f1 + f2
            f0, f1, f2 = f1, f2, ans
            yield ans


def transfer(ss: list[str]) -> str:
    ml = max([len(s) for s in ss])
    delta_map = {i: x for i, x in enumerate(fib(ml))}

    ans: list[str] = []

    for s in ss:
        new_s = []
        for idx, c in enumerate(s):
            delta = delta_map[idx]
            # 只需要走一轮
            # 先 %26 算出来是第几个字母，再加上 97 换算成对应字母
            asc = (ord(c) - 97 + delta) % 26 + 97
            new_s.append(chr(asc))

        ans.append("".join(new_s))

    return "\n".join(ans)


if __name__ == "__main__":
    print(transfer(["xy", "abc"]))
