package rocketpunch

import (
	"context"
	"encoding/json"
	"geekermeter-data/crawler"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var (
	baseURL = "https://www.rocketpunch.com/jobs?page=1"
)

func Rocketpunch() []crawler.Job {
	now := strconv.Itoa(time.Now().Year())
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
		log.Println("Current Page : ", i)
		temp := make([]crawler.Job, 200)
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
				temp[sliceCap].URL = "https://www.rocketpunch.com" + row.AttributeValue("href")
				temp[sliceCap].Origin = origin
				sliceCap++
			}
			sliceCap = tempSliceCap

			for _, row := range newbieNode {
				//int8 로 반환
				temp[sliceCap].Newbie = crawler.Getnewbie(row.Children[0].NodeValue)
				//crawler.Gentnewbie에서 목표값으로 반환
				//temp[sliceCap].Newbie = crawler.Newbie(crawler.Getnewbie(row.Children[0].NodeValue))
				log.Println(temp[sliceCap].Newbie)
				sliceCap++
			}
			sliceCap = tempSliceCap

			for _, row := range dateNode {
				//var tt string
				var kk []*cdp.Node
				p := strconv.Itoa(int(row.ChildNodeCount))
				_ = chromedp.Run(ctx,
					chromedp.Nodes(row.FullXPath()+"/span["+p+"]", &kk),
				)
				r, _ := regexp.Compile("등록")
				for _, pow := range kk {
					if r.MatchString(pow.Children[0].NodeValue) {
						temp[sliceCap].StartDate = now + crawler.ExceptKorean(pow.Children[0].NodeValue)
					}
				}
				sliceCap++
			}
		}
		temp = temp[0:sliceCap]
		crawledData = append(crawledData, temp...)
	}

	return crawledData
}

func BodyText(box crawler.Job, forname int) {
	ctx, cancel := chromedp.NewContext(context.Background())

	defer cancel()

	// run task list
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(box.URL),
		chromedp.Location(&loc),
	)
	crawler.ErrHandler(err)
	log.Println(box.URL)
	var nodes []*cdp.Node
	var clickNodes []*cdp.Node
	err = chromedp.Run(ctx,
		chromedp.Nodes("#wrap > div.eight.wide.job-content.column > section > h4 > a:nth-child(1)",
			&nodes))
	crawler.ErrHandler(err)

	clickctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err = chromedp.Run(clickctx,
		chromedp.Nodes("#wrap > div.eight.wide.job-content.column > section> div > a.see-more-text",
			&clickNodes),
	)

	for _, v := range clickNodes {
		err = chromedp.Run(ctx,
			chromedp.Click(v.FullXPath()),
		)
		crawler.ErrHandler(err)
	}

	var want string
	accepted := []string{"주요 업무", "업무 관련 기술 / 활동 분야", "채용 상세"}
	box.Content = make([]string, 3)
	count := 0
	for _, v := range nodes {
		err = chromedp.Run(ctx,
			chromedp.Text(v.FullXPath(), &want),
		)
		_, found := Find(accepted, want)
		if found == true {
			err = chromedp.Run(ctx,
				chromedp.Text(v.Parent.Parent.FullXPath(), &box.Content[count]),
			)
			crawler.ErrHandler(err)
			count++
		}
	}

	box.Content = box.Content[:count]
	toJson, _ := json.Marshal(box)
	t := strconv.Itoa(forname)
	_ = ioutil.WriteFile("./dataset/tmp/"+t+".json", toJson, 0644)

}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

//datetime은 크롤링의 대상이되는 날짜로 들어온다.
func Start(forname int, datetime string) int {
	log.Println("Start crawl Rocketpunch")
	list := Rocketpunch()
	log.Println("End crawl Rocketpunch")
	for _, row := range list {
		if datetime == row.StartDate {
			BodyText(row, forname)
			forname++
		}
	}
	return forname
}
