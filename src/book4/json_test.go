// json 探索
package book4

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
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

func TestJsonType(t *testing.T) {
	//data,err := json.Marshal(movies)	// 紧凑输出
	data,err := json.MarshalIndent(movies, "", "	")	// 格式化输出
	if err !=nil{
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n",data)
}

//////////////////// test github interface

const IssuesUrl = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
}