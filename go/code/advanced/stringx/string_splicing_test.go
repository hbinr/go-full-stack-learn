package stringx

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

/*
	字符串拼接：
	1.+=
	2.fmt.Sprint()    性能表现最差
	3.strings.join()  性能表现最好
	4.strings.Builder
	5.bytes.Buffer
*/

var languageStr = []string{"go", "java", "python", "c++", "hello"}

// BenchmarkStringSimple += 拼接
func BenchmarkStringSimple(b *testing.B) {
	// 5660358	       214 ns/op	      32 B/op	       5 allocs/op
	for i := 0; i < b.N; i++ {
		for _, v := range languageStr {
			v += "-"
		}
	}
}

// BenchmarkStringSprint  fmt.Sprint() 拼接
func BenchmarkStringSprint(b *testing.B) {
	//  1615084	       687 ns/op	     104 B/op	      10 allocs/op
	for i := 0; i < b.N; i++ {
		for _, v := range languageStr {
			fmt.Sprint(v)
		}
	}
}

// BenchmarkStringJoin strings.Join() 拼接
func BenchmarkStringJoin(b *testing.B) {
	// 13181829	        87.3 ns/op	      32 B/op	       1 allocs/op
	for i := 0; i < b.N; i++ {
		strings.Join(languageStr, "-")
	}
}

// BenchmarkStringBuffer bytes.Buffer 拼接
func BenchmarkStringBuffer(b *testing.B) {
	// 10167104	       117 ns/op	      64 B/op	       1 allocs/op
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		for _, v := range languageStr {
			b.WriteString(v)
			b.WriteString("-")
		}
	}
}

// BenchmarkStringBuilder strings.Builder 拼接
func BenchmarkStringBuilder(b *testing.B) {
	//  7824363	       150 ns/op	      56 B/op	       3 allocs/op
	for i := 0; i < b.N; i++ {
		var b strings.Builder
		for _, v := range languageStr {
			b.WriteString(v)
			b.WriteString("-")
		}
	}
}
