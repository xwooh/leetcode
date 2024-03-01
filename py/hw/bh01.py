"""
给定一个非空字符串S，其被N个‘-’分隔成N+1的子串，给定正整数K，要求除第一个子串外，其余的子串每K个字符组成新的子串，并用‘-’分隔。对于新组成的每一个子串，如果它含有的小写字母比大写字母多，则将这个子串的所有大写字母转换为小写字母；反之，如果它含有的大写字母比小写字母多，则将这个子串的所有小写字母转换为大写字母；大小写字母的数量相等时，不做转换。
输入描述:
输入为两行，第一行为参数K，第二行为字符串S。
输出描述:
输出转换后的字符串。
示例1
输入
3
12abc-abCABc-4aB@
输出
12abc-abc-ABC-4aB-@
说明
子串为12abc、abCABc、4aB@，第一个子串保留，后面的子串每3个字符一组为abC、ABc、4aB、@，abC中小写字母较多，转换为abc，ABc中大写字母较多，转换为ABC，4aB中大小写字母都为1个，不做转换，@中没有字母，连起来即12abc-abc-ABC-4aB-@
示例2
输入
12
12abc-abCABc-4aB@
输出
12abc-abCABc4aB@
说明
子串为12abc、abCABc、4aB@，第一个子串保留，后面的子串每12个字符一组为abCABc4aB@，这个子串中大小写字母都为4个，不做转换，连起来即12abc-abCABc4aB@

作者：yaozi
链接：https://leetcode.cn/circle/discuss/niKSMZ/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
"""


def cout(s: str) -> tuple[int, int]:
    small, big = 0, 0

    for c in s:
        if 97 <= ord(c) <= 122:
            small += 1
        elif 65 <= ord(c) <= 90:
            big += 1

    return small, big


def to_small(s: str) -> str:
    ws = []
    for c in s:
        if 65 <= ord(c) <= 90:
            ws.append(chr(ord(c) + 32))
        else:
            ws.append(c)
    return "".join(ws)


def to_big(s: str) -> str:
    ws = []
    for c in s:
        if 97 <= ord(c) <= 122:
            ws.append(chr(ord(c) - 32))
        else:
            ws.append(c)
    return "".join(ws)


def transfer(s: str, k: int) -> str:
    if "-" not in s:
        return s

    first, rest = s.split("-", maxsplit=1)
    rest = rest.replace("-", "")

    rws = [first]
    i = 0
    while i < len(rest):
        w = rest[i : i + k]
        i += k

        # 大小写转换
        sc, bc = cout(w)
        if sc > bc:
            # 小写
            w = to_small(w)
        elif bc > sc:
            # 大写
            w = to_big(w)
        rws.append(w)

    return "-".join(rws)


if __name__ == "__main__":
    assert transfer("12abc-abCABc-4aB@", 3) == "12abc-abc-ABC-4aB-@"
    assert transfer("12abc-abCABc-4aB@", 12) == "12abc-abCABc4aB@"
