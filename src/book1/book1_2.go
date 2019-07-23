// 1.2 命令行参数
package main

func main() {

	//fmt.Println("hello, world")

	// 第一种拼接参数的方法
	//var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = "/"
	//}
	//fmt.Println(s)

	// 第二种拼接参数的方法 range 返回值 第一个是索引，第二个是索引对应的参数值
	//s, sep := "", "/"
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//}
	//fmt.Println(s)

	// 使用String包的Join函数拼接
	//fmt.Println(strings.Join(os.Args[1:], " "))

	//fmt.Println(os.Args[1:])

	// 习题 1.1
	// 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
	//fmt.Println(os.Args[0])

	// 习题 1.2
	// 修改echo程序，使其打印每个参数的索引和值，每个一行。
	//for index, arg := range os.Args[1:] {
	//	fmt.Println(index, arg)
	//}

	// 习题1.3
	//  做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。
	//  （1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
	//startd:=time.Now()
	//s, sep := "", "/"
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//}
	//fmt.Println(s)
	//endd:=time.Now()
	//spendTime:= endd.Sub(startd)
	//fmt.Println("低效方式消耗时间：",spendTime)
	//
	//startg:=time.Now()
	//fmt.Println(strings.Join(os.Args[1:],"/"))
	//endg:=time.Now()
	//spendTimeg:= endg.Sub(startg)
	//fmt.Println("高效方式消耗时间：",spendTimeg)
}
