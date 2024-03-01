"""
题目
有一个特殊的五键键盘
    上面有A、Ctrl-C、Ctrl-X、Ctrl-V、Ctrl-A
    A键在屏幕上输出一个字母A
    Ctrl-C将当前所选的字母复制到剪贴板
    Ctrl-X将当前选择的字母复制到剪贴板并清空所选择的字母
    Ctrl-V将当前剪贴板的字母输出到屏幕
    Ctrl-A选择当前屏幕中所有字母

注意
1.剪贴板初始为空
2.新的内容复制到剪贴板会覆盖原有内容
3.当屏幕中没有字母时,Ctrl-A无效
4.当没有选择字母时Ctrl-C、Ctrl-X无效
5.当有字母被选择时A和Ctrl-V这两个输出功能的键,会先清空所选的字母再进行输出

要求
给定一系列键盘输入,输出最终屏幕上字母的数量。

输入描述:
输入为一行，为简化解析用数字12345分别代替A、Ctrl-C、Ctrl-X、Ctrl-V、Ctrl-A，输入的数字用空格分割

输出描述:
输出一个数字为屏幕上字母的总数量
"""


def output(inputs: list[int]) -> int:
    # 定义一个屏幕显示容器
    # 再定义一个剪贴板容器
    typed = 0
    is_selected = False
    pasted = 0

    for i in inputs:
        if i == 1:
            # a
            if is_selected:
                typed = 0
                is_selected = False
            typed += 1
        elif i == 2:
            # ctrl-c
            # 将选中的复制到剪贴板
            if is_selected:
                pasted = typed
        elif i == 3:
            # ctrl-x
            # 将选中的复制到剪贴板 & 清空已输入
            if is_selected:
                pasted = typed
                typed = 0
                is_selected = False
        elif i == 4:
            # ctrl-v
            # 把剪贴板内容输出
            if is_selected:
                typed = 0
                is_selected = False
            typed += pasted
        elif i == 5:
            # ctrl-a
            # 全选当前屏幕上的所有内容
            if typed:
                is_selected = True

    return typed


if __name__ == "__main__":
    print(output([1, 1, 1]))
    print(output([1, 1, 5, 1, 5, 2, 4, 4]))
