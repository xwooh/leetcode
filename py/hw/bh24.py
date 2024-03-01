"""
解密犯罪时间 全排列 dfs

题目描述

警察在侦破一个案件时，得到了线人给出的可能犯罪时间，形如“HH:MM”表示的时刻。根据警察和线人的约定，为了隐蔽，该时间是修改过的，解密规则为:利用当前出现过的数字，构造下一个距离当前时间最近的时刻，则该时间为可能的犯罪时间。每个出现数字都可以被无限次使用。

输入描述

形如HH:SS字符串，表示原始输入

输出描述

形如HH:SS的字符串，表示推理处理的犯罪时间。

备注

1.可以保证现任给定的字符串一定是合法的。

例如，“01:35”和“11:08”是合法的，“1:35”和“11:8”是不合法的

2.最近的时刻可能在第二天。

用例
20:12 => 20:20
因为用 2012 这几个数，排列得到的最近时间是 20:20
"""


def find_min(n: int, ns: list[int]) -> int:
    for i in ns:
        if i > n:
            return i
    return -1


def transfer(s: str) -> str:
    # 每一位都找比它大的第一个数 && 符合这一位的取值范围，如果没有取最小值

    nums = [int(s[0]), int(s[1]), int(s[-2]), int(s[-1])]
    snums = sorted(nums)
    mn = min(nums)

    for idx, n in enumerate(nums[::-1]):
        tidx = 3 - idx
        if idx == 0:
            x = find_min(n, snums)
            # print(f"{idx} {n} -> {x}, ", end="")
            if x == -1 or not 0 <= x <= 9:
                nums[tidx] = mn
                # print(f"use min {mn}")
            else:
                nums[tidx] = x
                # print(f"use x {x}")
                return f"{nums[0]}{nums[1]}:{nums[2]}{nums[3]}"

        elif idx == 1:
            x = find_min(n, snums)
            # print(f"{idx} {n} -> {x}, ", end="")
            if x == -1 or not 0 <= x <= 5:
                nums[tidx] = mn
                # print(f"use min {mn}")
            else:
                nums[tidx] = x
                # print(f"use x {x}")
                return f"{nums[0]}{nums[1]}:{nums[2]}{nums[3]}"

        elif idx == 2:
            x = find_min(n, snums)
            # print(f"{idx} {n} -> {x}, ", end="")
            if x == -1:
                nums[tidx] = mn
                # print(f"use min {mn}")
            else:
                if nums[3] == 2:
                    if not 0 <= x <= 3:
                        nums[tidx] = mn
                        # print(f"use min {mn}")
                    else:
                        nums[tidx] = x
                        # print(f"use x {x}")
                        return f"{nums[0]}{nums[1]}:{nums[2]}{nums[3]}"
                else:
                    if not 0 <= x <= 9:
                        nums[tidx] = mn
                        # print(f"use min {mn}")
                    else:
                        nums[tidx] = x
                        # print(f"use x {x}")
                        return f"{nums[0]}{nums[1]}:{nums[2]}{nums[3]}"

        elif idx == 3:
            x = find_min(n, nums)
            # print(f"{idx} {n} -> {x}, ", end="")
            if x == -1 or not 0 <= x <= 2:
                nums[idx] = mn
                # print(f"use min {mn}")
            else:
                nums[idx] = x
                # print(f"use x {x}")
                return f"{nums[0]}{nums[1]}:{nums[2]}{nums[3]}"

    return f"{nums[0]}{nums[1]}:{nums[2]}{nums[3]}"


if __name__ == "__main__":
    print(transfer("20:12"))
    print(transfer("23:59"))
    print(transfer("12:58"))
    print(transfer("18:52"))
    print(transfer("23:52"))
    print(transfer("09:17"))
    print(transfer("07:08"))
