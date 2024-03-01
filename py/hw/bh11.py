"""
按索引范围翻转文章片段 按单词下标区间翻转文章内容

题目描述

输入一个英文文章片段，翻转指定区间的单词顺序，标点符号和普通字母一样处理例如

输入字符串"I am a developer."，区间[0,3]，则输出"developer. a am I"

输入描述

使用换行隔开三个参数

第一个参数为英文文章内容即英文字符串

第二个参数为翻转起始单词下标(下标从 0 开始)

第三个参数为结束单词下标

输出描述

翻转后的英文文章片段所有单词之间以一个半角空格分隔进行输出
"""


def transfer(s: str, start: int, end: int) -> str:
    sl = s.split()

    i = max(0, start)
    j = min(end, len(sl) - 1)
    while i < j:
        sl[i], sl[j] = sl[j], sl[i]
        i += 1
        j -= 1

    return " ".join(sl)


if __name__ == "__main__":
    print(transfer("I am a developer.", 0, 3))
    print(transfer("I am a developer.", 1, 2))
