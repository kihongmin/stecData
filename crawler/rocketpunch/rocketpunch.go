package rocketpunch

import (
	"context"
	"geekermeter-data/crawler"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var (
	baseURL = "https://www.rocketpunch.com/jobs?page=1"
)

func Rocketpunch() []crawler.Job {
	var crawledData []crawler.Job
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var loc string
	err := chromedp.Run(ctx,
		chromedp.Location(&loc),
		chromedp.Navigate(baseURL),
	)
	crawler.ErrHandler(err)

	var totalPageNode []*cdp.Node
	var totalPage string
	//currPage := 0
	err = chromedp.Run(ctx,
		chromedp.Nodes("#search-results > div.ui.blank.right.floated.segment > div > div.tablet.computer.large.screen.widescreen.only > a:nth-child(7)",
			&totalPageNode, chromedp.ByID))
	crawler.ErrHandler(err)
	for _, row := range totalPageNode {
		totalPage = row.AttributeValue("data-query-add")[5:]
	}
	t, _ := strconv.Atoi(totalPage)
	for i := 1; i <= t; i++ { //페이지 단위
		temp := make([]crawler.Job, 100)
		var nodes []*cdp.Node
		var detailNode []*cdp.Node
		var origin string
		sliceCap := 0

		err := chromedp.Run(ctx, //해당 페이지의 company Node가
			chromedp.Navigate(`https://www.rocketpunch.com/jobs?page=`+strconv.Itoa(i)),
			chromedp.Sleep(2*time.Second),
			chromedp.Nodes("#company-list > div.company.item",
				&nodes, chromedp.ByQueryAll),
		)
		crawler.ErrHandler(err)
		for _, row := range nodes { //한 기업에서의 url 수집
			nodeNum := crawler.ExtractNum(row.PartialXPathByID())
			err = chromedp.Run(ctx,
				chromedp.Nodes("#company-list > div:nth-child("+nodeNum+") > div.content > div.company-jobs-detail > div.job-detail > div > a.nowrap.job-title.primary.link",
					&detailNode),
				chromedp.Text(`#company-list > div:nth-child(`+nodeNum+`) > div.content > div.company-name > a:nth-child(1) > h4 > strong`,
					&origin),
			)

			crawler.ErrHandler(err)

			for _, row := range detailNode {
				temp[sliceCap].Title = row.Children[0].NodeValue
				temp[sliceCap].URL = "https://www.rocketpunch.com/" + row.AttributeValue("href")
				temp[sliceCap].Origin = origin
				sliceCap++
			}
		}
		temp = temp[0:sliceCap]
		crawledData = append(crawledData, temp...)
	}
	return crawledData
}
