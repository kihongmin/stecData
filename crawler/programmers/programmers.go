package programmers

import (
	"context"
	"encoding/json"
	"geekermeter-data/crawler"
	"io/ioutil"
	"os"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var (
	baseURL = `https://programmers.co.kr/job`
)

func Programmers() []crawler.Job {
	var crawledData []crawler.Job
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(baseURL),
		chromedp.Location(&loc),
	)
	crawler.ErrHandler(err)

	for {
		temp := make([]crawler.Job, 20) // 최소단위 : 20

		//url & title node
		var nodes []*cdp.Node
		err = chromedp.Run(ctx, chromedp.Nodes("#list-positions-wrapper > ul > li > div.item-body > h4 > a", &nodes))
		crawler.ErrHandler(err)
		for i, row := range nodes {
			temp[i].URL = "https://programmers.co.kr/" + row.AttributeValue("href")
			temp[i].Title = row.Children[0].NodeValue
		}
		chromedp.Sleep(1 * time.Second)

		//company-name node
		var nameNode []*cdp.Node
		err = chromedp.Run(ctx, chromedp.Nodes("#list-positions-wrapper > ul > li> div.item-body > h5", &nameNode, chromedp.ByQueryAll))
		crawler.ErrHandler(err)
		for i, row := range nameNode {
			temp[i].Origin = row.Children[0].NodeValue
		}
		crawledData = append(crawledData, temp...)

		// move to next
		var checker []*cdp.Node

		chromedp.Sleep(2 * time.Second)
		err = chromedp.Run(ctx,
			chromedp.Nodes("#paginate > nav > ul > li.next.next_page.page-item > a", &checker),
		)
		crawler.ErrHandler(err)
		if checker[0].AttributeValue("rel") != "next" {
			break
		} else {
			err = chromedp.Run(ctx,
				chromedp.Click("#paginate > nav > ul > li.next.next_page.page-item > a", chromedp.NodeVisible),
			)
			crawler.ErrHandler(err)
		}
		chromedp.Sleep(1 * time.Second)
	}
	return crawledData
}

func BodyText(box crawler.Job) crawler.Job {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(box.URL),
		chromedp.Location(&loc),
	)
	crawler.ErrHandler(err)
	var position []*cdp.Node
	var requirements []*cdp.Node
	var preference []*cdp.Node
	err = chromedp.Run(ctx,
		chromedp.Sleep(2*time.Second),
		chromedp.Nodes("body > div.main > div.position-show > div > div > div.content-body.col-item.col-xs-12.col-sm-12.col-md-12.col-lg-8 > section.section-position > div > div > ul > li", &position, chromedp.ByQueryAll),
		chromedp.Nodes("body > div.main > div.position-show > div > div > div.content-body.col-item.col-xs-12.col-sm-12.col-md-12.col-lg-8 > section.section-requirements > div > div > ul > li", &requirements, chromedp.ByQueryAll),
		chromedp.Nodes("body > div.main > div.position-show > div > div > div.content-body.col-item.col-xs-12.col-sm-12.col-md-12.col-lg-8 > section.section-preference > div > div > ul > li", &preference, chromedp.ByQueryAll),
	)
	contentNum := 0
	box.Content = make([]string, 30)
	for _, row := range position {
		box.Content[contentNum] = row.Children[0].NodeValue
		contentNum++
	}
	for _, row := range requirements {
		box.Content[contentNum] = row.Children[0].NodeValue
		contentNum++
	}
	for _, row := range preference {
		box.Content[contentNum] = row.Children[0].NodeValue
		contentNum++
	}
	box.Content = box.Content[:contentNum]
	doc, _ := json.Marshal(box)
	err = ioutil.WriteFile("./articles.json", doc, os.FileMode(0644))
	return box
}
