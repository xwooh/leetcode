"""
求字符串中所有整数的最小和

题目描述

输入字符串s，输出s中包含所有整数的最小和。

说明:

字符串s，只包含a-zA-Z+-;合法的整数包括

1正整数一个或者多个0-9组成，如0230021022)负整数负号-开头，数字部分由一个

或者多个0-9组成，如-0 -012 -23 -00023

输入描述

包含数字的字符串

输出描述

所有整数的最小和
"""


def add(s: str) -> int:
    """负号开头的要一直往后取数"""
    ans = 0

    stack: list[str] = []
    neg = False
    for c in s:
        if c.isdigit():
            if neg:
                # 在累积负数
                stack.append(c)
            else:
                # 直接加上去
                ans += int(c)
        else:
            if neg and stack:
                # 把之前的负数处理掉
                ans -= int("".join(stack))
                stack = []

            neg = c == "-"

    if stack:
        ans -= int("".join(stack))

    return ans


if __name__ == "__main__":
    print(add("bb1234aa"))  #  10
    print(add("bb12-34aa"))  # -31
    print(add("b123-34-89a19aa-10"))  # -117
