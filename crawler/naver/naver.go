package naver

import (
	"context"
	"encoding/json"
	"fmt"
	"geekermeter-data/crawler"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
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
		chromedp.Click("#entType > a"),
		chromedp.Location(&loc),
	)
	crawler.ErrHandler(err)

	var nodes []*cdp.Node

	for i := 0; i < 7; i++ {
		k := rand.NormFloat64()*0.1 + 1
		time.Sleep(time.Duration(k) * time.Second)
		var check []*cdp.Node
		clickerr := chromedp.Run(ctx,
			chromedp.Nodes("#jobListDiv > ul >li>a", &nodes),
			chromedp.Nodes("#moreDiv", &check),
		)

		time.Sleep(time.Second * 2)

		err = chromedp.Run(ctx,
			chromedp.Click("#moreDiv > button"),
		)
		crawler.ErrHandler(clickerr)

	}

	chromedp.Sleep(3 * time.Second)
	log.Printf("\nclick success")

	//url node

	var titleNodes []*cdp.Node
	var dateNodes []*cdp.Node
	err = chromedp.Run(ctx,
		chromedp.Nodes("#jobListDiv > ul > li > a > span > strong",
			&titleNodes, chromedp.ByQueryAll),
		chromedp.Nodes("#jobListDiv > ul > li > a > span > em",
			&dateNodes, chromedp.ByQueryAll),
	)
	crawledData := make([]crawler.Job, len(titleNodes))

	for i, k := range nodes {
		crawledData[i].URL = "https://recruit.navercorp.com" + k.AttributeValue("href")
	}
	//title node
	for i, row := range titleNodes {
		crawledData[i].Title = row.Children[0].NodeValue

		crawledData[i].Origin = "naver"
	}

	for i, row := range dateNodes {
		crawledData[i].StartDate = crawler.Exceptspecial(row.Children[0].NodeValue[:10])
		log.Println(crawledData[i].StartDate)
	}

	return crawledData
}

func BodyText(box crawler.Job, forname int) {
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

	box.Content = make([]string, 1)
	doc.Find("#content > div > div.career_detail > div.dtl_context > div.context_area").Each(func(in int, tablehtml *goquery.Selection) {
		box.Content[0] = tablehtml.Text()
	})

	toJson, _ := json.Marshal(box)

	t := strconv.Itoa(forname)
	_ = ioutil.WriteFile("./dataset/tmp/"+t+".json", toJson, 0644)
	//_ = ioutil.WriteFile("./dataset/20200312/"+crawler.Exceptspecial(box.URL)+".json", toJson, 0644)

}
func Start(forname int) int {
	log.Println("Start crawl Naver")
	list := Naver()
	log.Println("End crawl Naver")
	for _, row := range list {
		BodyText(row, forname)
		forname++
	}
	return forname
}
