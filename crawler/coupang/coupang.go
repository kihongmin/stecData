package coupang

import (
	"context"
	"geekermeter-data/crawler"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var (
	baseURL = `https://rocketyourcareer.kr.coupang.com/%ea%b2%80%ec%83%89-%ec%a7%81%eb%ac%b4`
)

func Coupang() []crawler.Job {
	var crawledData []crawler.Job
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(baseURL),
		chromedp.Location(&loc),
	)

	crawler.ErrHandler(err)

	var totalPageNode []*cdp.Node
	var totalPage string
	err = chromedp.Run(ctx, chromedp.Nodes("#pagination-current-bottom", &totalPageNode, chromedp.ByID))
	crawler.ErrHandler(err)
	for _, row := range totalPageNode {
		//temp.url = "https://programmers.co.kr/" + row.AttributeValue("href")
		totalPage = row.AttributeValue("max")
	}
	t, _ := strconv.Atoi(totalPage)
	for i := 0; i < t; i++ {
		temp := make([]crawler.Job, 15) // 최소단위 : 10

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
		crawler.ErrHandler(err)
		for k, row := range nodes {
			temp[k].URL = "https://rocketyourcareer.kr.coupang.com" + row.AttributeValue("href")
			temp[k].Origin = "Coupang"
		}
		for k, row := range titleNode {
			//temp.url = "https://programmers.co.kr/" + row.AttributeValue("href")
			temp[k].Title = row.Children[0].NodeValue
		}
		crawledData = append(crawledData, temp...)
	}
	/*
		for _, dat := range crawledData {
			log.Printf("%s", dat.Title)
			log.Printf("%s", dat.URL)
			log.Printf("%s", dat.Origin)
		}
	*/
	return crawledData
}
