package main

import "fmt"
import "os"
import "io/ioutil"
import "net/http"
import "encoding/json"
import "strings"

type Result struct {
	Result  ResultCode
	Data    Data
}
type ResultCode struct {
	Msg string
	Code int
}
type Data struct {
	Entries []Mycore
	Query string
	Language string
	Type string
}
type Mycore struct {
	Explain string
	Entry string
}

var (
	BaseUrl = "https://dict.youdao.com/suggest?doctype=json&cache=false&le=en&q="
)

func main() {

	strs := os.Args
	strs = strs[1:]
	n := len(strs)
	if n <= 0 {
		fmt.Print("dict命令后最少跟一个参数 \n 用法： dict [单词] \n")
		return
	}
	q := strs[0]
	if m := len(q); m > 20{
		fmt.Print("dict命令参数最长为20 \n")
		return
	}
	defer func() {
		if err:=recover(); err != nil {
			fmt.Println(err)
		}
	}()
	resp, _ := http.Get(BaseUrl + q)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
    // if err != nil {
    //     fmt.Println("read from resp.Body failed,err:", err)
    //     return
    // }
	var res Result
	
	e := json.Unmarshal(body, &res)
	if e != nil {
		fmt.Println(e)
	}
	
	rnt := res.Data.Entries
	var bol bool = false
	for _, mc := range rnt {
		if strings.EqualFold(q, mc.Entry) {
			bol = true
			fmt.Println(mc.Explain)
		}
	}
	if !bol {
		fmt.Println("未找到准确匹配，以下是一些相近的词 ")
		for _, mc := range rnt {	
			fmt.Printf("%s  :  %s \n", mc.Entry, mc.Explain)
		}
	}
}
