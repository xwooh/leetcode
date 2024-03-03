"""
华为od机试 分积木 逻辑分析、按位异或运算

题目描述

Solo 和 koko 是两兄弟，妈妈给了他们一大堆积木，每块积木上都有自己的重量。
现在他们想要将这些积分成两堆。
哥哥 Solo 负责分配，
弟弟 koko 要求两个人获得的积木总重量"相等”(根据 Koko 的逻辑)，个数不同，不然就会哭，

但 koko 只会先将两个数转成二进制再进行加法，而且总会忘记进位(每个进位都忘记)如当 25 (11101) 加 11 (01011) 时，
koko 得到的计算结果是 18 (10010) :

  11001
+ 01011
--------
  10010

Solo 想要尽可能使自己得到的积木总重量最大，且不让 koko 哭

输入描述

第一行是一个整数 N(2≤N≤100)，表示有多少块积木;

第二行为空格分开的 N 个整数Ci(1≤Ci≤106)，表示第 i 块积木的重量

输出描述

如果能让 koko 不哭，输出 Solo 所能获得积木的最大总重量; 否则输出“NO”。
"""


def sp(weights: list[int]) -> int:
    """
    两个重要的概念
    1. 弟弟的不进位加法就是 => 异或运算
        1.1. 两个数一样，进行异或运算后就是 0
                所以要想分成两份重量一样的，必须所有数异或后为 0
    2. 如果所有的数异或运算后为 0，去掉任意一个数，剩下数和这个数是一样的二进制数
        因为这样异或后为 0
    """

    if len(weights) < 2:
        # 不足三个无法分
        return -1

    s = weights[0]
    for w in weights[1:]:
        s ^= w

    if s != 0:
        return -1

    weights.sort()
    return sum(weights[1:])


if __name__ == "__main__":
    print(sp([3, 5, 6]))