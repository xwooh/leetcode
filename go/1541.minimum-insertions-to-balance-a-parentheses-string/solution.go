// Created by Yuan at 2023/10/26 17:06
// leetgo: dev
// https://leetcode.cn/problems/minimum-insertions-to-balance-a-parentheses-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minInsertions(s string) (ans int) {

	var rn int

	for _, c := range s {
		if string(c) == "(" {
			if rn%2 == 1 {
				// 因为一个 ( 必定需要 2 个 ) 才能平衡，如果此时前面只有奇数个 )，
				//	那表明前面出现了一个 ) 去填补了 ) rn 的空缺，所以这边要再插入一个 ) 去补上缺的一个 )
				//	手动插入一个 ) 即 ans++，手动插入一个 ) 后对 rn 的需求对应也少一个，即 rn--
				ans++ // 插入一个 )
				rn--  // 上面插入了 )，那就对 ) 需求 -1
			}
			rn += 2
		} else if string(c) == ")" {
			rn -= 1

			if rn == -1 {
				// 多出来了一个 )
				ans++  // 需要插入一个 (
				rn = 1 // ) 的需求变为 1
			}
		}
	}

	ans += rn

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := minInsertions(s)

	fmt.Println("\noutput:", Serialize(ans))
}
