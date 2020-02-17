package programmers

import (
	"context"
	"geekermeter-data/crawler"
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

	// for _, dat := range crawledData {
	// 	log.Printf("%s", dat.Title)
	// 	log.Printf("%s", dat.URL)
	// 	log.Printf("%s", dat.Origin)
	// }

	// log.Println(count, cap(crawledData), len(crawledData))
	return crawledData
}
