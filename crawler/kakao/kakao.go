package kakao

import (
	"context"
	"geekermeter-data/crawler"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var (
	baseURL = "https://careers.kakao.com/jobs?company=ALL&keyword=&page=1"
)

func Kakao() []crawler.Job {
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
	err = chromedp.Run(ctx, chromedp.Nodes("#mArticle > div > div.paging_list > span > a.change_page.btn_lst", &totalPageNode, chromedp.ByID))
	crawler.ErrHandler(err)
	for _, row := range totalPageNode {
		//temp.url = "https://programmers.co.kr/" + row.AttributeValue("href")
		totalPage = row.AttributeValue("href")[11:]
	}
	t, _ := strconv.Atoi(totalPage)

	for i := 0; i < t; i++ {
		temp := make([]crawler.Job, 10)

		var nodes []*cdp.Node
		var titleNode []*cdp.Node
		var originNode []*cdp.Node
		if i == t-1 {
			err = chromedp.Run(ctx,
				chromedp.Sleep(2*time.Second),
				//url을 모두 node에 저장
				chromedp.Nodes("#mArticle > div > ul.list_notice > li > div > div > div > a", &nodes, chromedp.ByQueryAll),
				chromedp.Nodes("#mArticle > div > ul.list_notice > li > div > div > div > a > span", &titleNode, chromedp.ByQueryAll),
				chromedp.Nodes("#mArticle > div > ul.list_notice > li > div > div > span.field_front", &originNode, chromedp.ByQueryAll),
			)
		} else {
			err = chromedp.Run(ctx,
				chromedp.Sleep(2*time.Second),
				//url을 모두 node에 저장
				chromedp.Nodes("#mArticle > div > ul.list_notice > li > div > div > div > a", &nodes, chromedp.ByQueryAll),
				chromedp.Nodes("#mArticle > div > ul.list_notice > li > div > div > div > a > span", &titleNode, chromedp.ByQueryAll),
				chromedp.Nodes("#mArticle > div > ul.list_notice > li > div > div > span.field_front", &originNode, chromedp.ByQueryAll),
				chromedp.Click("#mArticle > div > div.paging_list > span > a:nth-child(14)", chromedp.NodeVisible),
				//다음 버튼 클릭
			)
		}

		crawler.ErrHandler(err)

		for k, row := range nodes {
			temp[k].URL = "https://careers.kakao.com" + row.AttributeValue("href")
		}
		for k, row := range titleNode {
			temp[k].Title = row.Children[0].NodeValue
		}
		for k, row := range originNode {
			temp[k].Origin = strings.TrimSpace(row.Children[0].NodeValue)
		}
		crawledData = append(crawledData, temp...)
	}
	return crawledData
}
