package book5

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println(sub(10, 2))
}

// 拥有2个int型参数和1个int型返回值的函数 举例四种
func add(x int, y int) int {
	return x + y
}
func add2(int,int) int {
	return 0
}
func add3(x int, _ int) int {
	return x
}
func sub(x int, y int) (z int) {
	z = x - y
	return
}
