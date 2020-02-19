package crafton

import (
	"context"
	"geekermeter-data/crawler"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var (
	baseURL = `https://krafton.recruiter.co.kr/app/jobnotice/list`
)

func Crafton() []crawler.Job {
	var crawledData []crawler.Job
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(baseURL),
		chromedp.Location(&loc),
	)
	crawler.ErrHandler(err)

	var totalPageNode []*cdp.Node
	var totalPage string
	err = chromedp.Run(ctx,
		chromedp.Nodes("#content > div.paging-wrapper.middle-set > div > button.btn.btn-paging.btn-small.btn-circle.fa.fa-angle-double-right", &totalPageNode, chromedp.ByID))
	crawler.ErrHandler(err)

	for _, row := range totalPageNode {
		//temp.url = "https://programmers.co.kr/" + row.AttributeValue("href")
		totalPage = row.AttributeValue("pageindex")
	}
	t, _ := strconv.Atoi(totalPage)

	for i := 2; i <= t+1; i++ {
		temp := make([]crawler.Job, 10) // 최소단위 : 10
		var nodes []*cdp.Node
		if i == t+1 {
			err = chromedp.Run(ctx,
				chromedp.Sleep(2*time.Second),
				chromedp.Nodes("#divJobnoticeList > ul > li > div > h2 > a", &nodes, chromedp.ByQueryAll),
			)
		} else {
			err = chromedp.Run(ctx,
				chromedp.Sleep(2*time.Second),
				chromedp.Nodes("#divJobnoticeList > ul > li > div > h2 > a", &nodes, chromedp.ByQueryAll),
				chromedp.Click("#content > div.paging-wrapper.middle-set > div > ul > li:nth-child("+strconv.Itoa(i)+") > a", chromedp.NodeVisible),
			)
		}

		crawler.ErrHandler(err)

		for k, n := range nodes {
			temp[k].URL = "https://krafton.recruiter.co.kr/app/jobnotice/view?systemKindCode=MRS2&jobnoticeSn=" + n.AttributeValue("data-jobnoticesn")
			temp[k].Title = n.Children[0].NodeValue
			temp[k].Origin = "Crafton"
		}
		crawledData = append(crawledData, temp...)
	}
	return crawledData
}
