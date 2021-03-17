/*
 * @Author: duanhaobin
 * @Date: 2021-03-17 17:38:55
 * @LastEditTime: 2021-03-17 18:46:29
 * @FilePath: \go-full-stack-learn\go\code\interview\stringx\02_reverse_byte_arr_test.go
 */

package stringx

import (
	"fmt"
	"testing"
)

/*
	问题：：
	编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
	不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

	示例 1：

	>输入：["h","e","l","l","o"]
	>输出：["o","l","l","e","h"]

	示例 2：

	>输入：["H","a","n","n","a","h"]
	>输出：["h","a","n","n","a","H"]

	分析：
	这是一道相当简单的经典题目，直接上题解：使用双指针进行反转字符串。

	假设输入字符串为`["h","e","l","l","0"]`

	- 定义left和right分别指向首元素和尾元素
	- 当`left < right` ，进行交换。
	- 交换完毕，`left++，right--`
	- 直至`left == right`

*/

func TestReverseByteArr(t *testing.T) {
	fmt.Println("res: ", string(reverseByteArr([]byte("hello"))))

}

func reverseByteArr(s []byte) []byte {
	right := len(s) - 1
	left := 0
	if left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}
