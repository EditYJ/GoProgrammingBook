//并发获取多个URL--一个小型的网速测试工具
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)		//创建通道
	for _,url := range os.Args[1:] {
		go fetch(url, ch)
	}
	// 主协程负责取数据
	// 当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，
	// 直到另一个goroutine往这个channel里写入、或者接收值，这样两个goroutine才会继续执行channel操作之后的逻辑。
	// 在这个例子中，每一个fetch函数在执行时都会往channel里发送一个值(ch <- expression)，主函数负责接收这些值(<-ch)。
	// 这个程序中我们用main函数来接收所有fetch函数传回的字符串，可以避免在goroutine异步执行还没有完成时main函数提前退出。
	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elepsed\n", time.Since(start).Seconds()) 	//计算总时间
}

// 访问url 并把相应的结果传入通道
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil{
		ch<- fmt.Sprint(err)
		return
	}
	// ioutil.Discard 可以把这个变量看作一个垃圾桶，可以向里面写一些不需要的数据
	// 主要是因为这里只需要[io.Copy()]返回的字节数，所以我们不关心[resp.Body]的具体内容
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil{
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch<- fmt.Sprintf("%.2fs %7d %s",secs, nbytes, url)
}