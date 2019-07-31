package book5

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"testing"
)

// [pre][post]是可选的，遍历每个节点的前后都会调用[pre]和[post]
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// 提取网页标题
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 检查返回的Content-Type是否为超文本格式("text/html;charset=utf-8")
	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/html" && !strings.HasPrefix(contentType, "text/html;") {
		return fmt.Errorf("%s 的类型为 %s, 不是需要的 text/html", url, contentType)
	}

	//进行html.Parse把html转换为*Node
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("转换%s的HTML出错，错误信息为: %v", url, err)
	}

	// 打印出title节点的内容
	// 此函数作为[forEachNode]的[pre]入参
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}

	// 递归遍历全部节点
	forEachNode(doc, visitNode, nil)

	return nil
}

// 捕获已知异常事例
func soleTitle(doc *html.Node) (title string, err error) {
	// 某个错误的特殊类型
	type bailout struct{}

	// 根据类型捕获异常
	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("错误：拥有多个title标签")
		default:
			panic(p)
		}
	}()

	//
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title !=""{
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)

	if title ==""{
		return "", fmt.Errorf("没有title标签")
	}
	return title, nil
}

func TestDeferred(t *testing.T) {
	_ = title("https://juejin.im/post/5d26914de51d45775746b9be")
}
