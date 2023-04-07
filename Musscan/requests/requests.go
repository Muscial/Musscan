package requests

//发包，匹配
import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var URL = flag.String("url", "", "input url")

type requsetsData struct { //请求包里的值
	Server string
	Header string
	Body   string
	Url    string
	Title  string
}

func Requsets(url string) *requsetsData {
	resp, err := http.Get(url)
	var re_data requsetsData
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()
	a, _ := ioutil.ReadAll(resp.Body)
	re_data.Body = string(a)
	Headerdecoded, _ := json.Marshal(resp.Header)
	re_data.Header = string(Headerdecoded)
	re_data.Url = url
	re_data.Server = resp.Header.Get("Server")
	return &re_data
}

// 匹配match 匹配两种写法，分开写和和其写，这里匹配写的死，大部分从chinfo师傅那cv的
func Check_title(checkfofa string, body string) bool { //checkfofa 是rules里面的content
	re := regexp.MustCompile("<title.*>(.*?)</title>")
	title := re.FindStringSubmatch(body)
	if title != nil {
		if title[1] == checkfofa {
			return true
		}

	}
	return false
}

func Check_Banner(checkfofa string, body string) bool { //匹配bannner
	re := regexp.MustCompile(`(?im)<\\s*banner.*>(.*?)<\\s*/\\s*banner>`)
	matchArr := re.FindAllStringSubmatch(body, -1)
	var i int
	for i = 0; i < len(matchArr); i++ {
		if checkfofa == matchArr[i][1] {
			return true
		}
	}
	return false
}

func Check_Header(checkfofa string, response *requsetsData) bool { //匹配header 有问题
	//header := response.Header
	//Headerdecoded, _ := json.Marshal(response.Header)
	Headerdecoded := response.Header
	result := strings.Contains(strings.ToLower(string(Headerdecoded)), strings.ToLower(checkfofa)) //fofa里面是小写
	return result
}

func Check_Body(checkfofa string, body string) bool {
	result := strings.Contains(strings.ToLower(body), strings.ToLower(checkfofa))
	return result
}

func Check_Server(checkfofa string, response *requsetsData) bool { //有没有一种可能Server也在相应头里面
	//Headerdecoded, _ := json.Marshal(response.Header)
	Headerdecoded := response.Header
	result := strings.Contains(strings.ToLower(string(Headerdecoded)), strings.ToLower(checkfofa)) //fofa里面是小写
	return result
}

func Check(match string, checkfofa string, body string, response *requsetsData) bool {
	result := false
	if match == "title_contains" {
		result = Check_title(checkfofa, body)
		return result
	} else if match == "banner_contains" { //go else if 要同一行
		result = Check_Banner(checkfofa, body)
		return result
	} else if match == "header_contains" {
		result = Check_Header(checkfofa, response)
		return result
	} else if match == "body_contains" {
		result = Check_Body(checkfofa, body)
		return result
	} else if match == "server_contains" {
		result = Check_Server(checkfofa, response)
		return result
	}
	return result
}
