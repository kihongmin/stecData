package rocketpunch

import (
	"context"
	"encoding/json"
	"fmt"
	"geekermeter-data/crawler"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
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
		var newbieNode []*cdp.Node
		var dateNode []*cdp.Node
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
				chromedp.Nodes("#company-list > div:nth-child("+nodeNum+") > div.content > div.company-jobs-detail > div.job-detail > div > span.job-stat-info",
					&newbieNode),
				chromedp.Nodes("#company-list > div:nth-child("+nodeNum+") > div.content > div.company-jobs-detail > div.job-detail > div.job-dates",
					&dateNode),
				chromedp.Text(`#company-list > div:nth-child(`+nodeNum+`) > div.content > div.company-name > a:nth-child(1) > h4 > strong`,
					&origin),
			)

			crawler.ErrHandler(err)
			tempSliceCap := sliceCap
			for _, row := range detailNode {
				temp[sliceCap].Title = row.Children[0].NodeValue
				temp[sliceCap].URL = "https://www.rocketpunch.com/" + row.AttributeValue("href")
				temp[sliceCap].Origin = origin
				sliceCap++
			}

			sliceCap = tempSliceCap
			for _, row := range newbieNode {
				temp[sliceCap].Newbie = crawler.OnlyKorean(row.Children[0].NodeValue)
				sliceCap++
			}

			sliceCap = tempSliceCap
			for _, row := range dateNode {
				var tt string
				p := strconv.Itoa(int(row.ChildNodeCount))
				err = chromedp.Run(ctx,
					chromedp.Text("#company-list > div:nth-child("+nodeNum+") > div.content > div.company-jobs-detail > div.job-detail > div.job-dates > span:nth-child("+p+")",
						&tt),
				)
				temp[sliceCap].StartDate = crawler.ExceptKorean(tt)
				sliceCap++
			}

		}
		temp = temp[0:sliceCap]
		crawledData = append(crawledData, temp...)
	}

	return crawledData
}

func BodyText(box crawler.Job) {
	res, err := http.Get(box.URL)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal()
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("No url found")
		log.Fatal(err)
	}

	accepted := []string{"주요 업무", "업무 관련 기술 / 활동 분야", "채용 상세"}
	box.Content = make([]string, 3)
	count := 0
	doc.Find("#wrap > div.eight.wide.job-content.column > section > h4").Each(func(in int, tablehtml *goquery.Selection) {
		_, found := Find(accepted, tablehtml.Text())
		if found == true {
			box.Content[count] = tablehtml.Parent().Text()
			count++
		}
	})
	box.Content = box.Content[:count]

	toJson, _ := json.Marshal(box)
	_ = ioutil.WriteFile("./dataset/new/"+crawler.Exceptspecial(box.URL)+".json", toJson, 0644)

}
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func Start() {
	list := Rocketpunch()
	for _, row := range list {
		BodyText(row)
	}
}
