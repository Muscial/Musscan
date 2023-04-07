package jsonRead

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// json的结构体
type CmsFeature []struct {
	RuleID         string `json:"rule_id"` //RuleID，类型为string，标签为json:"rule_id"，这个标签表示在JSON格式的数据中，这个成员的名称为rule_id。
	Level          string `json:"level"`
	Softhard       string `json:"softhard"`
	Product        string `json:"product"`
	Company        string `json:"company"`
	Category       string `json:"category"`
	ParentCategory string `json:"parent_category"`
	Rules          [][]struct {
		Match   string `json:"match"`
		Content string `json:"content"`
	} `json:"rules"`
}

// 读取文件
func Decode() CmsFeature {
	jsonFile, err1 := os.Open(`./fofa.json`) //打开json文件并进行错误处理
	if err1 != nil {
		panic(err1.Error())
	}
	fmt.Println("Have Opened fofa.jsonQAQ")
	defer jsonFile.Close()               //Go语言的 defer 语句会将其后面跟随的语句进行延迟处理
	byteValue, _ := io.ReadAll(jsonFile) //读取json数据
	var CMSList CmsFeature
	err2 := json.Unmarshal(byteValue, &CMSList)
	if err2 != nil {
		panic(err2.Error())
	}
	return CMSList
}

// 读取urllist
func UrlListRead(File string) []string {
	file, err := os.Open(File) //打开文件
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	readBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}
	results := strings.Split(string(readBytes), "\n")
	return results
}

/*
func ReadFile1(path string) error {
	fileHanle, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer fileHanle.Close()
	readBytes, err := ioutil.ReadAll(fileHanle)
	if err != nil {
		return err
	}
	results := strings.Split(string(readBytes), "\n")
	fmt.Printf("read result:%v", results)
	// 遍历每一行，并作为参数调用某个函数
	for _, line := range results {
		doSomething(line) // 这里你可以替换成你想要调用的函数
	}
	return nil
}

// 定义一个示例函数，打印出参数
func doSomething(arg string) {
	fmt.Println("The argument is:", arg)
}
*/
