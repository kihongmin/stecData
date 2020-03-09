package nexon

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
	baseURL = "https://career.nexon.com/user/recruit/notice/noticeList"
)

func Nexon() []crawler.Job {
	var crawledData []crawler.Job
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(baseURL),
		chromedp.Location(&loc),
	)
	crawler.ErrHandler(err)

	err = chromedp.Run(ctx,
		chromedp.Click("#container > ul > li:nth-child(1)", chromedp.NodeVisible))
	crawler.ErrHandler(err)

	var totalPageNode []*cdp.Node
	var totalPage string
	err = chromedp.Run(ctx,
		chromedp.Nodes("#con_right > div.content > div > a.last", &totalPageNode))
	crawler.ErrHandler(err)

	for _, row := range totalPageNode {
		totalPage = row.AttributeValue("href")[20:22]
	}
	t, _ := strconv.Atoi(totalPage)

	for i := 0; i <= t; i++ {
		var nodes []*cdp.Node
		var titleNode []*cdp.Node
		var newbieNode []*cdp.Node
		var countNode []*cdp.Node

		err := chromedp.Run(ctx,
			chromedp.Sleep(2*time.Second),
			chromedp.Nodes("#con_right > div.content > table > tbody > tr > td.tleft.fc_02 > a", &nodes, chromedp.ByQueryAll),
			chromedp.Nodes("#con_right > div.content > table > tbody > tr > td.tleft.fc_02 > a > span", &titleNode, chromedp.ByQueryAll),
			chromedp.Nodes("#con_right > div.content > table > tbody > tr > td:nth-child(2)", &newbieNode, chromedp.ByQueryAll),
			chromedp.Nodes("#con_right > div.content > table > tbody", &countNode, chromedp.ByQueryAll),
		)
		crawler.ErrHandler(err)
		var count int64
		for _, row := range countNode {
			count = row.ChildNodeCount
		}
		temp := make([]crawler.Job, count)

		for l, row := range nodes {
			temp[l].URL = "https://career.nexon.com" + row.AttributeValue("href")
		}
		for l, row := range titleNode {
			temp[l].Title = row.Children[0].NodeValue
			temp[l].Origin = "nexon"
		}
		for l, row := range newbieNode {
			temp[l].Newbie = row.Children[0].NodeValue
		}

		for i := 1; i <= int(count); i++ {
			t := strconv.Itoa(i)
			err := chromedp.Run(ctx,
				chromedp.Text("#con_right > div.content > table > tbody > tr:nth-child("+t+") > td:nth-child(6)", &temp[i-1].StartDate),
			)
			crawler.ErrHandler(err)
		}

		crawledData = append(crawledData, temp...)

		if i != t {
			err = chromedp.Run(ctx,
				chromedp.Click("#con_right > div.content > div > a.next", chromedp.NodeVisible),
			)
			crawler.ErrHandler(err)
		}
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
	box.Content = make([]string, 1)
	doc.Find("#con_right > div.content > div.list_txt01").Each(func(in int, tablehtml *goquery.Selection) {
		box.Content[0] = tablehtml.Text()
	})

	toJson, _ := json.Marshal(box)
	_ = ioutil.WriteFile("./dataset/new/"+crawler.Exceptspecial(box.URL)+".json", toJson, 0644)

}
