/*
 * @Author: duanhaobin
 * @Date: 2021-03-17 18:36:40
 * @LastEditTime: 2021-03-17 18:57:48
 * @FilePath: \go-full-stack-learn\go\code\interview\stringx\03_last_string_char_test.go
 */
package stringx

import (
	"fmt"
	"strings"
	"testing"
)

/*
	题目：给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1 。

	案例：
		s = "leetcode"
		返回 0.

		s = "loveleetcode",
		返回 2.
	分析：
	方法一：
		Go标准库有 strings包，Index()返回当前元素的索引，LastIndex返回当前元素最后一次出现的索引

		比较两个索引是否相等即可，相等则满足条件，跳出循环，否则继续遍历

	方法二：
		由于字母共有 26 个，所以我们可以声明一个 26 个长度的数组（该种方法在本类题型很常用）因为字符串中字母可能是重复的，
		所以我们可以先进行第一次遍历，在数组中记录每个字母的最后一次出现的所在索引。然后再通过一次循环，比较各个字母第一次出现的索引是否为最后一次的索引。
		如果是，我们就找到了我们的目标，如果不是我们将其设为 -1（标示该元素非目标元素）如果第二次遍历最终没有找到目标，直接返回 -1即可。
*/

func TestLastStringChar(t *testing.T) {
	fmt.Println("res: ", lastStringChar("loveleetcode"))
	fmt.Println("res2: ", lastStringChar2("loveleetcode"))
}

// 方法一：调用 标准库 strings 的方法
func lastStringChar(s string) int {
	idx := -1
	l := len(s)
	if l == 0 {
		return -1
	}
	for i := 0; i < l; i++ {
		curIdx := strings.IndexByte(s, s[i])
		lastIdx := strings.LastIndexByte(s, s[i])
		if curIdx == lastIdx {
			idx = curIdx
			break
		}
	}
	return idx
}

// 方法一：调用 标准库 strings 的方法
func lastStringChar2(s string) int {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
