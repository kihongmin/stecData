package naver

import (
	"context"
	"geekermeter-data/crawler"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var (
	baseURL = "https://recruit.navercorp.com/naver/job/list/developer"
)

func Naver() []crawler.Job { //아직 개발 직군만 크롤링임.

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(baseURL),
		chromedp.Location(&loc),
	)
	crawler.ErrHandler(err)
	PageCount := 10
	for {
		chromedp.Sleep(2 * time.Second)
		clickerr := chromedp.Run(ctx,
			chromedp.Click("#moreDiv > button"),
		)
		PageCount += 10
		if clickerr != nil {
			break
		}
	}
	chromedp.Sleep(2 * time.Second)
	log.Printf("\nclick success")
	crawledData := make([]crawler.Job, PageCount)
	//url node
	var nodes []*cdp.Node
	var titleNodes []*cdp.Node
	err = chromedp.Run(ctx,
		chromedp.Nodes("#jobListDiv > ul > li > a", &nodes),
		chromedp.Nodes("#jobListDiv > ul > li > a > span > strong", &titleNodes, chromedp.ByQueryAll),
	)

	for i, k := range nodes {
		crawledData[i].URL = "https://recruit.navercorp.com" + k.AttributeValue("href")
	}
	//title node
	for i, row := range titleNodes {
		log.Println(i)
		crawledData[i].Title = row.Children[0].NodeValue
		crawledData[i].Origin = "naver"
	}

	//date node
	/*
		var dateNode []*cdp.Node
		if err = chromedp.Run(ctx, chromedp.Nodes("#jobListDiv > ul > li > a > span > em", &dateNode, chromedp.ByQueryAll)); err != nil {
			log.Fatal(err)
		}
		for i, row := range dateNode {
			crawledData[i].date = row.Children[0].NodeValue
		}*/
	/*
		for _, dat := range crawledData {
			log.Printf("%s", dat.Title)
			log.Printf("%s", dat.URL)
			log.Printf("%s", dat.Origin)
		}
	*/
	return crawledData
}
