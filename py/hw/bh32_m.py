"""
找最小数 栈结构

题目描述

给一个正整数NUM1，计算出新正整数NUM2，NUM2为NUM1中移除N位数字后的结果，需要使得NUM2的值最小。

输入描述

输入的第一行为一个字符串，字符串由0-9字符组成，记录正整数NUM1，NUM1长度小于322.
输入的第二行为需要移除的数字的个数，小于NUM1长度。

输出描述

输出一个数字字符串，记录最小值NUM2。

用例
输入:
    2615371
    4
输出:
    131
"""


def transfer(n1: str, k: str) -> str:
    # 单调栈，从前往后拿
    # 另外要计算好剔除了多少个元素，当剔除够了就不要再剔除了

    if not n1.isdigit() or not k.isdigit():
        return ""

    rm = int(k)
    if len(n1) <= rm:
        return "0"

    stack: list[int] = []  # 非递减栈
    i = 0
    while i < len(n1):
        # 不断拿元素压栈
        while rm > 0 and stack and stack[-1] > int(n1[i]):
            stack.pop()
            rm -= 1
        stack.append(int(n1[i]))
        i += 1

    # 可能没剔除够
    while rm > 0:
        stack.pop()
        rm -= 1

    # 把头部 0 元素删掉
    while stack and stack[0] == 0:
        stack = stack[1:]

    if stack:
        return "".join([str(x) for x in stack])

    return "0"


if __name__ == "__main__":
    print(transfer("1071", "1"))  # 71
    print(transfer("2615371", "4"))  # 131
    print(transfer("8123981361872", "10"))  # 111
    print(transfer("8123981361872", "11"))  # 11
    print(transfer("8123981361872", "9"))  # 1112
