package coupang

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type data struct {
	url   string
	title string
}

func Coupang() {
	var crawledData []data
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://rocketyourcareer.kr.coupang.com/%ea%b2%80%ec%83%89-%ec%a7%81%eb%ac%b4`),
		//chromedp.Click("#container > ul > li:nth-child(1)", chromedp.NodeVisible),
		// 인재채용 페이지까지 들어옴
		chromedp.Location(&loc),
		//chromedp.Click("#page-container > div > div.signin-wrapper > form > div.clearfix > button",chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	var totalPageNode []*cdp.Node
	var totalPage string
	if err := chromedp.Run(ctx, chromedp.Nodes("#pagination-current-bottom", &totalPageNode, chromedp.ByID)); err != nil {
		log.Fatal(err)
	}
	for _, row := range totalPageNode {
		//temp.url = "https://programmers.co.kr/" + row.AttributeValue("href")
		totalPage = row.AttributeValue("max")
	}
	t, _ := strconv.Atoi(totalPage)
	for i := 0; i < t; i++ {
		temp := make([]data, 15) // 최소단위 : 10

		var nodes []*cdp.Node
		var titleNode []*cdp.Node

		err := chromedp.Run(ctx,
			chromedp.Sleep(2*time.Second),
			//url을 모두 node에 저장
			chromedp.Nodes("#search-results-list > ul > li > a", &nodes, chromedp.ByQueryAll),
			chromedp.Nodes("#search-results-list > ul > li > a > h3", &titleNode, chromedp.ByQueryAll),
			chromedp.Click("#pagination-bottom > div.pagination-paging > a.next", chromedp.NodeVisible),
			//다음 버튼 클릭
		)
		if err != nil {
			log.Fatal(err)
		}
		for i, row := range nodes {
			temp[i].url = "https://rocketyourcareer.kr.coupang.com" + row.AttributeValue("href")
		}
		for i, row := range titleNode {
			//temp.url = "https://programmers.co.kr/" + row.AttributeValue("href")
			temp[i].title = row.Children[0].NodeValue
		}
		crawledData = append(crawledData, temp...)
	}

	for _, dat := range crawledData {
		log.Printf("%s", dat.title)
		log.Printf("%s", dat.url)
	}
}
