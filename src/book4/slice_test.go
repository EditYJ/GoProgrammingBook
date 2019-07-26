package book4

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	a := [...]string{"sads", "asd", "Fsdwe", "gwe", "rwe", "vsc"}
	b := [...]string{"sads", "asd", "Fsdwe", "gwe", "rwe", "vsc"}
	//reverse(a[:])
	fmt.Println(equal(a[:], b[:]))
}

// 反转切片
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//比较两个切片之间是否相同
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// append函数的理解
func TestAppendFunc(t *testing.T) {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = append(x, 1, 2, 3, 4)
		fmt.Printf("%d\tcap=%d\t%v\n", x, cap(y), y)
		// 开始x容量是零，所以第一次append函数就给y分配了新空间，
		// x的指向还是那个容量为零的空间，所以没有这个等于的话，x永远指向那个容量为零的空间
		x = y
	}
}

// append 函数的影子 appendInt
func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	// IF|如果新切片的长度小于等于被增加切片x的容量
	// z则为x所指向的内存空间
	// ELSE|
	// 首先求z的容量[zcap],首先让[zcap]等于xlen+newlen
	// 如果[zcap]小于于两倍的x长度，则扩展[zcap]的大小为2 * xlen
	// 此时新返回的切片z则指向了另一段新的切片空间，和x所指向的内存空间不一样了
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
