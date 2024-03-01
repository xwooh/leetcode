"""
https://zhuanlan.zhihu.com/p/641740464
这个链接通过栈来解析

二叉树中序遍历，栈结构

题目描述

根据给定的二叉树结构描述字符串，输出该二叉树按照中序遍历结果字符串。
中序遍历顺序为:左子树，根结点，右子树。

输入描述

由大小写字母、左右大括号、逗号组成的字符串:
    字母代表一个节点值，左右括号内包含该节点的子节点，左右子节点使用逗号分隔，
    逗号前为空则表示左子节点为空，没有逗号则表示右子节点为空。

如：a{b{d,e{g,h{,I}}},c{f}}

又树节点数最大不超过100。

注:输入字符串格式是正确的，无需考虑格式错误的情况。

输出描述

输出一个字符串为二叉树中序遍历各节点值的拼接结果
"""

from dataclasses import dataclass
from typing import Optional


@dataclass
class Node:
    val: str = ""
    left: Optional["Node"] = None
    right: Optional["Node"] = None


def build(s: str) -> Node | None:
    # 获取左右子节点信息
    if s == "":
        return None
    elif len(s) == 1:
        return Node(s)
    elif len(s) == 3:
        return Node(s[0])
    elif len(s) == 4:
        return Node(s[0], left=Node(s[2]))

    root = Node(s[0])

    i = 2
    if s[i] == ",":
        # 没有左节点
        l = ""
        r = s[i + 1 : -1]
    elif s[i + 1] == ",":
        l = s[i]
        r = s[i + 2 : -1]
    else:
        start = i
        i += 1
        c = 1
        while c:
            i += 1
            if s[i] == "{":
                c += 1
            elif s[i] == "}":
                c -= 1
        l = s[start : i + 1]
        # 判断有没有右节点
        if i < len(s) - 1 and s[i + 1] == ",":
            r = s[i + 2 : -1]
        else:
            r = ""

    ln = build(l)
    rn = build(r)
    root.left = ln
    root.right = rn

    return root


def travel(n: Node | None, c: list[str]):
    if n is None:
        return

    c.append(n.val)
    travel(n.left, c)
    travel(n.right, c)


def output(s: str) -> str:
    n = build(s)
    ans = []
    travel(n, ans)
    return "".join(ans)


if __name__ == "__main__":
    print(output("a{b{d,e{g,h{,I}}},c{f}}"))
