package programmers

import (
	"context"
	"encoding/json"
	"fmt"
	"geekermeter-data/crawler"
	"io/ioutil"
	"log"
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
			temp[i].URL = "https://programmers.co.kr" + row.AttributeValue("href")
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

func BodyText(box crawler.Job) { //현재 쓸데없는 값까지 하는 중->예외처리 실패로 인해..
	log.Println(box.URL)
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(box.URL),
		chromedp.Location(&loc),
	)
	crawler.ErrHandler(err)

	var nodes []*cdp.Node
	var sectionName string
	err = chromedp.Run(ctx,
		chromedp.Nodes("body > div.main > div.position-show > div > div > div.content-body.col-item.col-xs-12.col-sm-12.col-md-12.col-lg-8>section",
			&nodes))
	crawler.ErrHandler(err)

	acceptList := []string{"section-stacks", "section-position", "section-requirements", "section-preference"}
	presentList := make([]string, 4)
	presentNum := 0
	for _, row := range nodes {
		sectionName = row.AttributeValue("class")
		_, found := Find(acceptList, sectionName)
		if found == true {
			presentList[presentNum] = sectionName
			presentNum++
		}
	}
	presentList = presentList[:presentNum]

	box.Content = make([]string, presentNum)
	for index, name := range presentList {
		selurl := "body > div.main > div.position-show > div > div > div.content-body.col-item.col-xs-12.col-sm-12.col-md-12.col-lg-8 > section." + name
		box.Content[index] = getdetail(ctx, selurl)
	}

	doc, _ := json.Marshal(box)
	_ = ioutil.WriteFile("./dataset/new/"+crawler.Exceptspecial(box.URL)+".json", doc, 0644)

}

func getdetail(ctx context.Context, selector string) string {
	var target string
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
			return
		}
	}()

	if err := chromedp.Run(ctx,
		chromedp.Text(selector, &target)); err != nil {
		panic(err)
	}

	return target
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
