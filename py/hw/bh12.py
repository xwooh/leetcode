"""
最大嵌套括号深度
题目描述：
现有一字符串仅由 ‘(’， ‘)’， ‘{’， ‘}’， ‘[’， ']'六种括号组成。 若字符串满足以下条件之一，则为无效字符串：

1. 任一类型的左右括号数量不相等；
2. 存在未按正确顺序（先左后右）闭合的括号。
输出括号的最大嵌套深度，若字符串无效则输出 0。 0≤字符串长度≤100000

输入描述:
一个只包括 ‘(’， ‘)’， ‘{’， ‘}’， ‘[’， ']'的字符串

输出描述：
一个整数，最大的括号深度

示例 1：

输入
[]
输出
1
说明
有效字符串，最大嵌套深度为1

示例 2：

输入
([]{()})
输出
3
说明
有效字符串，最大嵌套深度为3

示例 3：

输入
(]
输出
0
说明
无效字符串，有两种类型的左右括号数量不相等
"""


def matched(s: str, c: str) -> bool:
    if (s == "(" and c == ")") or (s == "[" and c == "]") or (s == "{" and c == "}"):
        return True
    return False


def calc(s: str) -> int:
    # 从左往右压栈
    # 每加入一个新的左括号，先认为匹配，ans = max(ans, len(stack))
    # 匹配上右括号，那就栈顶出栈
    # 完全匹配上就是 ans，匹配不上就返回0

    if len(s) % 2 != 0:
        return 0

    ans = 0
    stack = []
    for c in s:
        if not stack:
            if c in {")", "]", "}"}:
                # 右括号打头就不对了
                return 0
            else:
                stack.append(c)
                ans = max(ans, len(stack))
        else:
            if c in {")", "]", "}"}:
                if not matched(stack[-1], c):
                    # 右括号不匹配
                    return 0
                else:
                    # 右括号匹配，弹出
                    stack = stack[: len(stack) - 1]
            else:
                # 新的左括号
                stack.append(c)
                ans = max(ans, len(stack))

    return ans


if __name__ == "__main__":
    print(calc("[]") == 1)
    print(calc("[}") == 0)
    print(calc("([()]{([])})[({[()]})([{}])]") == 5)
    print(calc("([]{{()}})") == 4)
    print(calc("([{[{}]}]{{()}})") == 5)
    print(calc("([{}]({[]}){{()}}){{[]}}") == 4)
