// 3.5 字符串
package book3

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringType(t *testing.T) {
	//result := baseNameNew("D:\\Work\\周报\\俞杰工作周报 -20190519.xlsx")
	result := dealStrFloat("12321341232132.4321242")
	fmt.Println(result)
}

// 取出文件名的函数
func basename(s string) string {
	// 去除‘/’前面的字符串
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '\\' {
			s = s[i+1:]
			break
		}
	}
	// 去除‘.’后面的字符串
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// 改进版,利用 [strings.LastIndex()]
func baseNameNew(s string) string {
	i := strings.LastIndex(s, string('\\'))
	s = s[i+1:]
	if dot := strings.LastIndex(s, string('.')); dot > 0 {
		s = s[:dot]
	}
	return s
}

//将一个表示整数值的字符串，每隔三个字符插入一个逗号分隔符，例如“12345”处理后成为“12,345”。
func formatStrNumber(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return formatStrNumber(s[:n-3]) + "," + s[n-3:]
}

func formatStrNumberR(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return s[:3] + "," + formatStrNumberR(s[3:])
}

func dealStrFloat(s string) string {
	n:=strings.LastIndex(s, string('.'))
	res:=formatStrNumber(s[:n])+"."+formatStrNumberR(s[n+1:])
	return res
}
