package eval

import "testing"

// 接口Expr来表示Go语言中任意的表达式
type Expr interface{}

// Var类型表示对一个变量的引用
type Var string

// literal类型表示一个浮点型常量
type literal float64

// unary和binary类型表示有一到两个运算对象的运算符表达式
type unary struct {
	op rune
	x  Expr
}
type binary struct {
	op   rune
	x, y Expr
}

// call类型表示对一个函数的调用, 我们限制它的fn字段只能是pow，sin或者sqrt。
type call struct {
	fn   string
	args []Expr
}

type Env map[Var]float64

func TestEvalProgram(t *testing.T) {

}
