// 1.3 查找重复的行
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// 标准输入
	//counts := make(map[string]int)
	//input := bufio.NewScanner(os.Stdin)
	//for input.Scan() {
	//	if input.Text() == "end" {
	//		break
	//	}
	//	counts[input.Text()]++
	//}
	//
	//for key, value := range counts {
	//	if value > 1 {
	//		fmt.Printf("%d\t%s\n", value, key)
	//	}
	//}

	// 文件输入
	//  修改dup2，出现重复的行时打印文件名称。
	//	counts := make(map[string]int)
	//	files := os.Args[1:]	// 读取路径参数
	//	if len(files) == 0 {	// 如果无路径参数，则标准输入
	//		countsLines(os.Stdin, counts)
	//	} else {
	//		for _, arg := range files {
	//			f, err := os.Open(arg)
	//			if err != nil {
	//				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	//				continue
	//			}
	//			countsLines(f, counts)
	//
	//
	//
	//			f.Close()
	//		}
	//	}
	//
	//	for line, n := range counts {
	//		if n > 1 {
	//			fmt.Printf("%d\t%s\n", n, line)
	//		}
	//	}
	//}
	//
	//func countsLines(file *os.File, counts map[string]int) {
	//	input := bufio.NewScanner(file)
	//	for input.Scan() {
	//		if input.Text() =="end"{
	//			break
	//		}
	//		counts[input.Text()]++
	//		if counts[input.Text()]>1 {
	//			fmt.Println(file.Name())
	//		}
	//	}

	// 文件输入 取消接收标准输入
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
