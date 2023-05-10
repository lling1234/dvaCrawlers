package main

import (
	"dvaCrawlers/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {

	jsonData := body2json(model.BodyData)
	fmt.Println("jsonData", string(jsonData))

	ListItems := make([]model.ListItem, 0)
	json.Unmarshal(jsonData, &ListItems)

	for _, v := range ListItems {
		fmt.Println("v.Title", v.Title)
		fmt.Println("v.Href", v.Href)
		fmt.Println("v.Time", v.Time)
	}

}
func httpGet() {
	url := "https://www.19fzw.com/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP请求出错：", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容出错：", err)
		return
	}

	err2 := ioutil.WriteFile("bady.txt", []byte(body), 0644)
	if err2 != nil {
		panic(err)
	}

	fmt.Println("body", string(body))
	re := regexp.MustCompile(`<div class="new-list-page">([\s\S]*?)<div class="down-list">
    <div class="title">下载专区`)
	match := re.FindStringSubmatch(string(body))
	fmt.Println("match")
	fmt.Println(match)

	if len(match) > 1 {
		content := match[1]
		fmt.Println(match)
		err2 := ioutil.WriteFile("newlistpage.txt", []byte(content), 0644)
		if err2 != nil {
			panic(err)
		}
	} else {
		fmt.Println("没有匹配到内容")
	}
}

func body2json(body string) []byte {
	re := regexp.MustCompile(`<li[^>]*>\s*<a href="([^"]+)"[^>]+>([^<]+)</a>\s*<font[^>]*>([^<]+)</font>\s*</li>`)
	matches := re.FindAllStringSubmatch(body, -1)

	data := make([]map[string]string, 0, len(matches))

	for _, match := range matches {
		if len(match) > 3 {
			data = append(data, map[string]string{
				"href":  match[1],
				"title": match[2],
				"time":  match[3],
			})
		}
	}

	jsonData, err := json.Marshal(data)
	if err == nil {
		fmt.Println(string(jsonData))
		return jsonData
	} else {
		fmt.Println("无效的输入")
		return nil
	}
}
