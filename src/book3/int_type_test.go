// 3.1 整型
package book3

import (
	"fmt"
	"testing"
)

// 位运算
func TestBitOperation (t *testing.T) {
	var x uint8 = 1<<1 | 1<<5 // 10 | 100000 = 0010 0010
	var y uint8 = 1<<1 | 1<<2 // 10 | 100 = 0000 0110

	fmt.Printf("x: %08b\n", x)
	fmt.Printf("y: %08b\n", y)

	fmt.Printf("x&y: %08b\n", x&y)   // 与 100010 & 110 = 100010
	fmt.Printf("x|y: %08b\n", x|y)   // 或 100010 | 110 = 100110
	fmt.Printf("x^y: %08b\n", x^y)   // 异或(不一样就为1) 100010 ^ 110 = 100100
	fmt.Printf("x&^y: %08b\n", x&^y) // 位清空 100010 &^ 110 = 100000

	for i := uint(0); i < 8; i++ {
		//fmt.Print(i)
		if x&(1<<i) != 0{
			fmt.Println(i)
		}
	}
}
