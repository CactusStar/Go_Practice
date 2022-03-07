package main

import (
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// type Swoop struct {
// 	url string
// 	header map[string]string
// }

func main() {
	url := "https://www.google.com/?hl=zh_CN"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36")
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	name := doc.Find("a.gb_f").Text()
	fmt.Println(name)
	// 设定关闭响应体
	defer resp.Body.Close()
	// 读取响应体
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(body))

	fmt.Scanf("h")
}
