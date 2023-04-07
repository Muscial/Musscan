package main

import (
	"Musscan/checkFinger"
	"Musscan/jsonRead"
	"flag"
	"fmt"
	"strings"
)

func main() {
	Banner()
	var Url string = ""
	var File string = ""
	finger_data := jsonRead.Decode() //解析fofa.json

	flag.StringVar(&Url, "Url", "", "input a url")
	flag.StringVar(&File, "file", "", "input path to your url.txt")
	if !flag.Parsed() { //解析传参
		flag.Parse()
	}
	if Url != "" { //单个url
		fmt.Println("------------------------------------------------------")
		fmt.Println(Url)
		checkFinger.Tocheck(Url, finger_data) //调用函数首个字母大写
		fmt.Println("------------------------------------------------------")

	}
	if File != "" { //多url传参
		urlbyline := jsonRead.UrlListRead(File)
		for _, line := range urlbyline {
			fmt.Println("------------------------------------------------------")
			fmt.Println(line)
			line = strings.TrimSpace(line)         //去掉逆天换行符
			checkFinger.Tocheck(line, finger_data) // 这里你可以替换成你想要调用的函数
			fmt.Println("------------------------------------------------------")
		}

	}
	//go run main.go -Url https://1.117.175.65/
	//go run main.go -file urllist.txt
	//checkFinger.Tocheck(*url, finger_data)

	//发包， 匹配指纹，打印指纹
	//fmt.Println(finger_data)
	//data := requests.Requsets("http://www.o2takuxx.com")
	//fmt.Println(data.Header)

}
func Banner() {
	banner := `
     ___           ___           ___           ___           ___           ___           ___     
    /\__\         /\__\         /\  \         /\  \         /\  \         /\  \         /\__\    
   /::|  |       /:/  /        /::\  \       /::\  \       /::\  \       /::\  \       /::|  |   
  /:|:|  |      /:/  /        /:/\ \  \     /:/\ \  \     /:/\:\  \     /:/\:\  \     /:|:|  |   
 /:/|:|__|__   /:/  /  ___   _\:\~\ \  \   _\:\~\ \  \   /:/  \:\  \   /::\~\:\  \   /:/|:|  |__ 
/:/ |::::\__\ /:/__/  /\__\ /\ \:\ \ \__\ /\ \:\ \ \__\ /:/__/ \:\__\ /:/\:\ \:\__\ /:/ |:| /\__\
\/__/~~/:/  / \:\  \ /:/  / \:\ \:\ \/__/ \:\ \:\ \/__/ \:\  \  \/__/ \/__\:\/:/  / \/__|:|/:/  /
	  /:/  /   \:\  /:/  /   \:\ \:\__\    \:\ \:\__\    \:\  \            \::/  /      |:/:/  / 
	 /:/  /     \:\/:/  /     \:\/:/  /     \:\/:/  /     \:\  \           /:/  /       |::/  /  
	/:/  /       \::/  /       \::/  /       \::/  /       \:\__\         /:/  /        /:/  /   
	\/__/         \/__/         \/__/         \/__/         \/__/         \/__/         \/__/    

	`
	print(banner)
}
