package netmarble

import (
	"context"
	"log"
	"time"

	"geekermeter-data/crawler"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var (
	baseURL      = "https://m.netmarble.com/rem/www/noticelist.jsp"
	baseSelector = "#contents > div > div > div > div.recruit_list_wrapper > ul > li >"
)

func Netmarble() []crawler.Job {
	var crawledData []crawler.Job

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(baseURL),
		chromedp.Location(&loc),
	)
	crawler.ErrHandler(err)

	//find max page
	var pageNode []*cdp.Node
	var totalPage string
	clickerr := chromedp.Run(ctx,
		chromedp.Nodes("#pageCount", &pageNode, chromedp.ByID),
	)
	crawler.ErrHandler(clickerr)
	for _, row := range pageNode {
		totalPage = row.Children[0].NodeValue
	}

	for {
		temp := make([]crawler.Job, 10) // 최소단위 : 10

		chromedp.Sleep(2 * time.Second)
		//url
		var nodes []*cdp.Node
		var titleNode []*cdp.Node
		var levelNode []*cdp.Node
		var dateNode []*cdp.Node
		var groupNode []*cdp.Node

		err := chromedp.Run(ctx,
			chromedp.Nodes(baseSelector+"div.cw_jopinfo > a", &nodes),
			chromedp.Nodes(baseSelector+"div.cw_jopinfo > a > span.cw_title", &titleNode, chromedp.ByQueryAll),
			chromedp.Nodes(baseSelector+"div.cw_jopinfo > a > span.cw_info > span.cw_type", &levelNode, chromedp.ByQueryAll),
			chromedp.Nodes(baseSelector+"div.cw_group", &groupNode, chromedp.ByQueryAll),
			chromedp.Nodes(baseSelector+"div.cw_jopinfo > a > span.cw_info > span.cw_range", &dateNode, chromedp.ByQueryAll),
		)
		crawler.ErrHandler(err)

		for i, row := range nodes {
			temp[i].URL = "https://m.netmarble.com/rem/www" + row.AttributeValue("href")[1:]
		}
		//title
		for i, row := range titleNode {
			temp[i].Title = row.Children[0].NodeValue
		}
		//level
		/*
			for i, row := range levelNode {
				temp[i].level = row.Children[0].NodeValue
			}
		*/
		//group
		/*
			for i, row := range groupNode {
				temp[i].group = row.Children[0].NodeValue
			}
		*/
		//date
		/*
			for i, row := range dateNode {
				temp[i].date = row.Children[0].NodeValue
			}
		*/
		crawledData = append(crawledData, temp...)

		var currPage string
		chromedp.Sleep(2 * time.Second)
		clickerr := chromedp.Run(ctx,
			chromedp.Text("#contents > div > div > div > div.recruit_list_wrapper > div.recruit_pagination > span.page_current", &currPage),
		)
		if clickerr != nil {
			log.Fatal(clickerr)
		}

		if currPage == totalPage {
			break
		} else {
			err = chromedp.Run(ctx,
				chromedp.Click("#contents > div > div > div > div.recruit_list_wrapper > div.recruit_pagination > button.page_next", chromedp.NodeVisible),
			)

		}

		//fmt.Println(i)
	}
	for _, dat := range crawledData {
		log.Printf("%s", dat.Title)
		log.Printf("%s", dat.URL)
	}

	return crawledData
}
