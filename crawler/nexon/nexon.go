package nexon

import (
	"context"
	"geekermeter-data/crawler"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var (
	baseURL = "https://career.nexon.com/user/recruit/notice/noticeList"
)

func Nexon() []crawler.Job {
	var crawledData []crawler.Job
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(baseURL),
		chromedp.Location(&loc),
	)
	crawler.ErrHandler(err)

	err = chromedp.Run(ctx,
		chromedp.Click("#container > ul > li:nth-child(1)", chromedp.NodeVisible))
	crawler.ErrHandler(err)

	var totalPageNode []*cdp.Node
	var totalPage string
	err = chromedp.Run(ctx,
		chromedp.Nodes("#con_right > div.content > div > a.last", &totalPageNode))
	crawler.ErrHandler(err)

	for _, row := range totalPageNode {
		totalPage = row.AttributeValue("href")[20:22]
	}
	t, _ := strconv.Atoi(totalPage)

	for i := 0; i <= t; i++ {
		temp := make([]crawler.Job, 10)
		var nodes []*cdp.Node
		var titleNode []*cdp.Node

		err := chromedp.Run(ctx,
			chromedp.Sleep(2*time.Second),
			chromedp.Nodes("#con_right > div.content > table > tbody > tr > td.tleft.fc_02 > a", &nodes, chromedp.ByQueryAll),
			chromedp.Nodes("#con_right > div.content > table > tbody > tr > td.tleft.fc_02 > a > span", &titleNode, chromedp.ByQueryAll),
		)
		crawler.ErrHandler(err)

		for l, row := range nodes {
			temp[l].URL = "https://career.nexon.com" + row.AttributeValue("href")
		}
		for l, row := range titleNode {
			temp[l].Title = row.Children[0].NodeValue
			temp[l].Origin = "nexon"
		}

		crawledData = append(crawledData, temp...)

		if i != t {
			err = chromedp.Run(ctx,
				chromedp.Click("#con_right > div.content > div > a.next", chromedp.NodeVisible),
			)
			crawler.ErrHandler(err)
		}
	}

	for i, _ := range crawledData {
		log.Println(crawledData[i].URL)
		log.Println(crawledData[i].Title)
		log.Println(crawledData[i].Origin)
	}
	return crawledData
}
