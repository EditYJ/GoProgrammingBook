package book5

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func add1(r rune) rune {
	return r + 1
}

//// [pre][post]是可选的，遍历每个节点的前后都会调用[pre]和[post]
//func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
//	if pre != nil{
//		pre(n)
//	}
//	for c := n.FirstChild; c != nil; c = c.NextSibling {
//		forEachNode(c, pre, post)
//	}
//	if post!=nil{
//		post(n)
//	}
//}

var depth int

// 利用fmt.Printf的一个小技巧控制输出的缩进。
// %*s中的*会在字符串之前填充一些空格。
// 在例子中，每次输出会先填充depth*2数量的空格，再输出""，最后再输出HTML标签。
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

// strings.Map对字符串中的每个字符调用add1函数，
// 并将每个add1函数的返回值组成一个新的字符串返回给调用者。
func TestFuncValue(t *testing.T) {
	fmt.Println(strings.Map(add1, "HAL-9000"))
	node, err := html.Parse(Fetch("https://juejin.im/post/5d26914de51d45775746b9be"))
	if err!=nil{
		fmt.Printf("findlinks: %v\n", err)
	}
	forEachNode(node, startElement, endElement)
}
