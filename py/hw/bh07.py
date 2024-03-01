"""
给定一个正整数数组 检查数组中是否存在满足规则的数组组合 规则： A=B+2C

输入描述 第一行输出数组的元素个数 接下来一行输出所有数组元素 用空格隔开

输出描述 如果存在满足要求的数 在同一行里依次输出 规则里 A/B/C的取值 用空格隔开 如果不存在输出 []

示例1： 输入 4 2 7 3 0 输出 [7, 3, 2] 说明： 7=3+2*2

示例2： 输入 3 1 1 1 输出 [] 说明找不到满足条件的组合

作者：隔壁黄同学
链接：https://www.nowcoder.com/discuss/353150624449634304
来源：牛客网
"""


def find(nums: list[int]) -> list[int]:
    if len(nums) < 3:
        return []

    ans = []

    nums = sorted(nums)
    nl = len(nums)

    seen = set()

    i, j = 0, nl - 1
    for i in range(nl - 1):
        for j in range(nl - 1, 0, -1):
            for k in range(i + 1, j):
                if nums[j] == nums[i] + 2 * nums[k]:
                    if (nums[j], nums[i]) not in seen:
                        ans.append([nums[j], nums[i], nums[k]])
                        seen.add((nums[j], nums[i]))
                elif nums[j] == nums[k] + 2 * nums[i]:
                    if (nums[j], nums[k]) not in seen:
                        ans.append([nums[j], nums[k], nums[i]])
                        seen.add((nums[j], nums[k]))

    return ans


if __name__ == "__main__":
    print(find([5, 2, 3, 4, 7, 1, 1]))
    print(find([2, 4, 0]))
    print(find([1, 3, 2]))
    print(find([1, 0, 0, 5]))
