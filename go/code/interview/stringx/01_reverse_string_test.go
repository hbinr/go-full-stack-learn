/*
 * @Author: duanhaobin
 * @Date: 2021-03-17 17:21:59
 * @LastEditTime: 2021-03-17 18:46:18
 * @FilePath: \go-full-stack-learn\go\code\interview\stringx\01_reverse_string_test.go
 */
package stringx

import (
	"fmt"
	"testing"
)

/*
	问题:
	请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
	给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

	分析:
	翻转字符串其实是将一个字符串以中间字符为轴，前后翻转，即将str[len]赋值给str[0],将str[0] 赋值 str[len]。
*/
func TestReverseString(t *testing.T) {
	var res string
	res, _ = reverseString("123")
	fmt.Println(res)

	res, _ = reverseString("Hello World 中国")
	fmt.Println(res)
}

func reverseString(s string) (string, bool) {
	str := []rune(s) // 因GO默认使用UTF-8编码，所以对于len(中文),返回的字节长度为3，不符合字符串长度的逻辑
	// 因此使用将参数转换为rune类型，rune是int32的别名，占4个字节，使用Unicode编码，这样就可以对中文计算字符串长度了
	l := len(str)
	if l > 5000 {
		return s, false
	}
	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]

	}
	return string(str), true
}
