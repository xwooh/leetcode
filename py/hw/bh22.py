"""
仿LISP运算 栈结构

题目描述

LISP语言唯一的语法就是括号要配对

形如(OP P1 P2 ...)，括号内元素由单个空格分割。

其中第一个元素OP为操作符，后续元素均为其参数，参数个数取决于操作符类型。

注意:参数P1,P2也有可能是另外一个嵌套的(OP P1 P2...)，
当前OP类型为add/sub/mul/div(全小写) 分别代表整数的加减乘除法，简单起见，所有OP参数个数均为2。

举例:

输入:(mul 3 -7)输出:-21

输入:(add 1 2)输出:3

输入:(sub (mul 2 4) (div 9 3))输出:5

输入:(div 1 0)输出:error

- 题目涉及数字均为整数，可能为负
- 不考虑32位溢出翻转，计算过程中也不会发生32位溢出翻转除零错误时，输出“error”，
- 除法遇除不尽，向下取整，即3/2=1

输入描述

输入为长度不超过512的字符串，用例保证了无语法错误

输出描述

输出计算结果或者“error
"""


def do(op_stack: list[str], num_stack: list[int]):
    # print(f"{op_stack} {num_stack} => ", end="")
    op = op_stack.pop()

    b = num_stack.pop()
    a = num_stack.pop()

    if op == "add":
        c = a + b
    elif op == "sub":
        c = a - b
    elif op == "mul":
        c = a * b
    elif op == "div":
        c = a // b
    else:
        raise ValueError("语法错误")

    num_stack.append(c)
    # print(f"{op_stack} {num_stack}")


def cacl(s: str) -> str:
    """
    操作符压入 op 栈，数字压入 num 栈
    遇到右括号，弹出 op 栈顶元素
    从 num 栈弹出 2 个数字，用弹出的操作符进行计算，把结果再压入 num 栈
    """
    op_stack: list[str] = []
    num_stack: list[int] = []

    tmp = []

    i = 0
    while i < len(s):
        c = s[i]
        if not c.isdigit():
            if tmp:
                num_stack.append(int("".join(tmp)))
                tmp = []

        if c == ")":
            # 计算
            try:
                do(op_stack, num_stack)
            except:
                return "error"
            i += 1
        elif c == "a":
            op_stack.append("add")
            i += 3
        elif c == "s":
            op_stack.append("sub")
            i += 3
        elif c == "m":
            op_stack.append("mul")
            i += 3
        elif c == "d":
            op_stack.append("div")
            i += 3
        elif c.isdigit():
            tmp.append(c)
            i += 1
        else:
            i += 1

    return str(num_stack[-1])


if __name__ == "__main__":
    print(cacl("(div 12 (sub 45 45))"))
    print(cacl("(add(sub(mul 2 4)3)(div 9 3))"))
