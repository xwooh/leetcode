"""
数组拼接

题目描述

现在有多组整数数组，需要将它们合并成一个新的数组。

合并规则:
    从每个数组里按顺序取出固定长度的内容合并到新的数组中，取完的内容会删除掉，
    如果该行不足固定长度或者已经为空，则直接取出剩余部分的内容放到新的数组中，继续下一行。

输入描述

第一行是每次读取的固定长度，0<长度<10

第二行是整数数组的数目，0<数目<1000

第3-n行是需要合并的数组，不同的数组用回车换行分隔，数组内部用逗号分隔，最大不超过100个元素。

输出描述

输出一个新的数组，用逗号分隔
"""


def transfer(k: int, ns: str) -> str:
    nums_list = [
        [int(n.strip()) for n in nl.strip().split(",")] for nl in ns.strip().split()
    ]

    ans: list[str] = []
    max_len = max([len(x) for x in nums_list])
    i = 0

    while i < max_len:
        for nl in nums_list:
            if i >= len(nl):
                continue
            ans.extend(list(map(str, nl[i : i + k])))
        i += k

    return ",".join(ans)


if __name__ == "__main__":
    s = """2,5,6,7,9,5,7
        1,7,4,3,4"""
    print(transfer(3, s))
