// 微型web服务
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)
//，服务器每一次接收请求处理时都会另起一个goroutine，
// 这样服务器就可以同一时间处理多个请求。然而在并发情况下，
// 假如真的有两个请求同一时刻去更新count，那么这个值可能并不会被正确地增加；
// 所以此处使用了[mu.Lock()]和[mu.Unlock()]
func main() {
	http.HandleFunc("/", handler) //每个请求都会经过此handler,比较类似中间件(过滤层)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("192.168.8.42:8000", nil))
}

func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	_, _ = fmt.Fprintf(writer, "主页访问次数：%d\n", count)
	mu.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 代码示例1
	//mu.Lock()
	//count++
	//fmt.Println("有人访问！！！")
	//_, _ = fmt.Fprintf(w, "壮壮牛逼了,居然访问了：%q\n 欢迎欢迎 :)", r.URL.Path)
	//mu.Unlock()

	// 代码示例2-打印请求的各种信息
	// [r.Method]			访问方式(GET/POST/...)
	// [r.URL]				访问的URL路径
	// [r.Proto]			使用的HTTP协议标准
	// [r.Header]			请求头信息类型（貌似是个数组）
	// [r.Host]				主机地址
	// [r.RemoteAddr]	客户端的IP地址
	//_, _ = fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	//for k, v := range r.Header {
	//	_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	//}
	//_, _ = fmt.Fprintf(w, "Host = %q\n", r.Host)
	//_, _ = fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	//if err := r.ParseForm(); err != nil {
	//	log.Print(err)
	//}
	//for k, v := range r.Form {
	//	fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	//}

	// 例三 输出1.4中的GIF图
	lissajous(w)

}



func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

