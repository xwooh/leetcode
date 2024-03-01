"""
题目描述

RSA加密算法在网络安全世界中无处不在，它利用了极大整数因数分解的困难度，数据越大，安全系数越高，
给定一个32位正整数，请对其进行因数分解，找出是哪两个素数的乘积。

输入描述

个正整数num0<num<2147483647

输出描述

如果成功找到，以单个空格分割，从小到大输出两个素数，分解失败，请输出-1,-1
"""


def isPrime(n: int):
    if n <= 1:
        return False

    i = 2
    while i * i <= n:
        if n % i == 0:
            # 素数只能被 1 和 自己 整除
            return False

        i += 1

    return True


def quick_check(num: int):
    i = 2
    while i * i <= num:
        if isPrime(i) and num % i == 0 and isPrime(num // i):
            return i, num // i
        i += 1

    return -1, -1


if __name__ == "__main__":
    print(quick_check(15))
    print(quick_check(27))
    print(quick_check(5827387))
