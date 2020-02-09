package main

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func Crawler() {
	url := "https://careers.kakao.com/jobs?company=ALL&keyword=&page="

	for i := 1; i <= 19; i++ { //일단은 페이지 고정 -> 추후에 자동으로 확인하는 것 추가
		err := GetUrl(url + strconv.Itoa(i))
		if err == -1 {
			break
		}
	}
}

func GetUrl(url string) int {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	doc.Find("ul.list_notice li").Each(func(_ int, li *goquery.Selection) {
		TargetUrl, ok := li.Find("a").Attr("href")
		//title := li.Find("span.txt_tit").Text()
		if ok {
			fmt.Println("https://careers.kakao.com/" + TargetUrl)
			//fmt.Println("title")
		}
	})
	return 0
}

func main() {
	Crawler()
}
