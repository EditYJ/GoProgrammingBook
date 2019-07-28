// json 探索
package book4

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`	// omitempty(省略空值)选项，表示当Go语言结构体成员为空或零值时不生成该JSON对象（这里false为零值）
	Actors []string
}

var movies = []Movie{
	{
		Title: "Casablanca", Year: 1942, Color: false,
		Actors:[]string{"Humphrey Bogart", "Ingrid Bergman"},
	},	{
		Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors:[]string{"Paul Newman"},
	},	{
		Title: "Bullitt", Year: 1968, Color: true,
		Actors:[]string{"Steve McQueen", "JacqueLine Bisset"},
	},
}

var titles []struct{Title string}

func TestJsonType(t *testing.T) {
	data,err := json.Marshal(movies)	// 紧凑输出
	//data,err := json.MarshalIndent(movies, "", "	")	// 格式化输出
	err = json.Unmarshal(data, &titles)	// json转对象
	if err !=nil{
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n",titles)
}

//////////////////// test github interface

const IssuesUrl = "https://api.github.com/search/issues"

// 问题的搜索结果
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

// 每个问题的实体
type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreateAt time.Time `json:"created_at"`
	Body string
}

// 提出问题的用户
type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	//fmt.Println(strings.Join(terms, " "))

	resp, err := http.Get(IssuesUrl + "?q=" +q)
	if err !=nil{
		return nil,err
	}

	for k, v := range resp.Header {
		fmt.Printf("Header[%q] = %q\n",k, v)
	}
	// 如果请求失败 我们则必须关闭resp.Body
	if resp.StatusCode != http.StatusOK{
		resp.Body.Close()
		return nil, fmt.Errorf("搜索失败: %s",resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result);err != nil{
		resp.Body.Close()
		return nil,err
	}

	resp.Body.Close()
	return &result, nil
}

func TestGitHubInterface(t *testing.T) {
	result,err := SearchIssues([]string{"golang","flutter"})
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	for _,item := range result.Items{
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User, item.Title)
	}
}