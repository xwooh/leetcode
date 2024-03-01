"""
链接：https://www.nowcoder.com/questionTerminal/5382ff24fbf34a858b15f93e2bd85307
来源：牛客网

给定两个字符串 s和 t ，判断 s是否为 t 的子序列。
你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度n ~= 500,000），而 s 是个短字符串（长度 <=100）。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

进阶：时间复杂度 O(n) ，空间复杂度 O(n)

输入描述:
共两行，第一行为字符串s,  第二行为字符串t


字符串t的长度 1<=n<=500000


字符串s的长度 1<=m<=100



输出描述:
输出true或者是false，true表示是s是t的子序列，false表示s不是t的子序列
"""


class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        i, j = 0, 0

        while j < len(t):
            if i >= len(s):
                return True

            if s[i] == t[j]:
                i += 1
            j += 1

        return i >= len(s)
