// 测试递归
package book5

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

// 访问输入的网址，将其内容转换为字节流
func fetch(url string) *bytes.Reader {
	resp, err := http.Get(url)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	return bytes.NewReader(b)
}

// [html.Parse]转换输出*Node，此函数接受node 遍历全html，找出内部所有链接
func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// 检查html所有检点，输出每一个父节点的深度情况
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func TestVisit(t *testing.T) {
	doc, err := html.Parse(fetch("https://juejin.im/post/5d26914de51d45775746b9be"))
	if err != nil {
		fmt.Printf("findlinks: %v\n", err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func TestOutline(t *testing.T) {
	doc, err := html.Parse(fetch("https://juejin.im/post/5d26914de51d45775746b9be"))
	if err != nil {
		fmt.Printf("findlinks: %v\n", err)
	}
	outline(nil, doc)
}

// 测试fetch
func TestFetch(t *testing.T) {
	fmt.Println(fetch("https://juejin.im/post/5d26914de51d45775746b9be"))
}
