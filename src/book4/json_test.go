// json 探索
package book4

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	//"text/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"` // omitempty(省略空值)选项，表示当Go语言结构体成员为空或零值时不生成该JSON对象（这里false为零值）
	Actors []string
}

var movies = []Movie{
	{
		Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	}, {
		Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"},
	}, {
		Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "JacqueLine Bisset"},
	},
}

var titles []struct{ Title string }

func TestJsonType(t *testing.T) {
	data, err := json.Marshal(movies) // 紧凑输出
	//data,err := json.MarshalIndent(movies, "", "	")	// 格式化输出
	err = json.Unmarshal(data, &titles) // json转对象
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", titles)
}

//////////////////// test github interface

const IssuesUrl = "https://api.github.com/search/issues"

// 问题的搜索结果
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// 每个问题的实体
type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	Body     string
}

// 提出问题的用户
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// 普通模板
const temple = `{{.TotalCount}} issues:
{{range .Items}}
-----------------------------------
issue编号:	#{{.Number}}
创建用户:	{{.User.Login}}
标题:			{{.Title | printf "%.64s"}}
存在时间:	{{.CreateAt | getAgoToNowLength}}
{{end}}
`

// html模板
const htmlTemple = `<h1>{{.TotalCount}} issues</h1>
<table border="1">
<tr style='text-align: left'>
  <th>#</th>
  <th>状态</th>
  <th>作者</th>
  <th>标题</th>
  <th>存在时间</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td>{{.CreateAt | getAgoToNowLength}}</td>
</tr>
{{end}}
</table>
`

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	//fmt.Println(strings.Join(terms, " "))

	resp, err := http.Get(IssuesUrl + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// 打印返回的头部信息
	//for k, v := range resp.Header {
	//	fmt.Printf("Header[%q] = %q\n", k, v)
	//}

	// 如果请求失败 我们则必须关闭resp.Body
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("搜索失败: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

// 计算相差天数
// 输入小时数
// 返回距今多少天
func getAgoToNowLength(t time.Time) string {
	timeHoursLag := int(time.Since(t).Hours())
	var timeSpend bytes.Buffer
	if timeHoursLag <= 0 { //如果玩小于等于0 说明是在今天
		timeSpend.WriteString(strconv.Itoa(0))
		timeSpend.WriteString("天")
	} else { //如果差值大于0  说玩家的天数相差的起码有一天之前上
		day := timeHoursLag / 24     //s算出来之间相差多少天
		if (timeHoursLag % 24) > 0 { //还有剩余的消失 说明是 还有一天 就那就加上
			day = day + 1
		}
		switch {
		case day < 30:
			timeSpend.WriteString(strconv.Itoa(day))
			timeSpend.WriteString("天")
		case day >= 30 && day < 365:
			timeSpend.WriteString(strconv.Itoa(day/30))
			timeSpend.WriteString("月")
		case day >= 365:
			timeSpend.WriteString(strconv.Itoa(day/365))
			timeSpend.WriteString("年")
		}
	}
	return timeSpend.String()
}

func TestGitHubInterface(t *testing.T) {
	result, err := SearchIssues([]string{"golang", "flutter"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		//计算天数
		day := getAgoToNowLength(item.CreateAt)
		//fmt.Println(day," 天前")
		fmt.Printf("#%-5d %-5s前提问 %9.9s %.55s\n", item.Number, day, item.User, item.Title)
	}
}

func TestGitHubInterfaceWithTemple(t *testing.T) {
	// 1.template.New先创建并返回一个模板；
	// 2.Funcs方法将daysAgo等自定义函数注册到模板中，并返回模板；
	// 3.最后调用Parse函数分析模板。
	//result, err := template.New("report").
	//	Funcs(template.FuncMap{"getAgoToNowLength": getAgoToNowLength}).
	//	Parse(temple)
	//if err !=nil{
	//	log.Fatal(err)
	//}

	// 使用[template.Must]辅助函数处理模板转换可能出现的错误
	resultTemple := template.Must(template.New("report").
		Funcs(template.FuncMap{"getAgoToNowLength": getAgoToNowLength}).
		Parse(temple))

	// 从GitHub得到Issue关键字搜索结果
	result, err := SearchIssues([]string{"golang", "flutter"})
	if err != nil {
		log.Fatal(err)
	}
	if err := resultTemple.Execute(os.Stdout, result); err !=nil{
		log.Fatal(err)
	}
}

// 起web服务来返回查询到的相关内容
func TestWebServerDisplay(t *testing.T) {
	http.HandleFunc("/", indexHandle)
	log.Fatal(http.ListenAndServe("localhost:9768", nil))
}

// 处理主路由
func indexHandle(writer http.ResponseWriter, request *http.Request) {
	// 使用[template.Must]辅助函数处理模板转换可能出现的错误
	resultTemple := template.Must(template.New("report").
		Funcs(template.FuncMap{"getAgoToNowLength": getAgoToNowLength}).
		Parse(htmlTemple))

	// 从GitHub得到Issue关键字搜索结果
	result, err := SearchIssues([]string{"golang"})
	if err != nil {
		log.Fatal(err)
	}
	// 输入数据实体[result]
	resultTemple.Execute(writer, result);
}

