package netmarble

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"geekermeter-data/crawler"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var (
	baseURL      = "https://m.netmarble.com/rem/www/noticelist.jsp"
	baseSelector = "#contents > div > div > div > div.recruit_list_wrapper > ul > li >"
)

func Netmarble() []crawler.Job {
	var crawledData []crawler.Job

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(baseURL),
		chromedp.Location(&loc),
	)
	crawler.ErrHandler(err)

	//find max page
	var pageNode []*cdp.Node
	var totalPage string
	clickerr := chromedp.Run(ctx,
		chromedp.Nodes("#pageCount", &pageNode, chromedp.ByID),
	)
	crawler.ErrHandler(clickerr)
	for _, row := range pageNode {
		totalPage = row.Children[0].NodeValue
	}

	for {
		temp := make([]crawler.Job, 10) // 최소단위 : 10

		chromedp.Sleep(2 * time.Second)
		//url
		var nodes []*cdp.Node
		var titleNode []*cdp.Node
		var newbieNode []*cdp.Node
		var dateNode []*cdp.Node

		err := chromedp.Run(ctx,
			chromedp.Nodes(baseSelector+"div.cw_jopinfo > a", &nodes),
			chromedp.Nodes(baseSelector+"div.cw_jopinfo > a > span.cw_title", &titleNode, chromedp.ByQueryAll),
			chromedp.Nodes(baseSelector+"div.cw_jopinfo > a > span.cw_info > span.cw_type", &newbieNode, chromedp.ByQueryAll),
			chromedp.Nodes(baseSelector+"div.cw_jopinfo > a > span.cw_info > span.cw_range", &dateNode, chromedp.ByQueryAll),
		)
		crawler.ErrHandler(err)

		for i, row := range nodes {
			temp[i].URL = "https://m.netmarble.com/rem/www" + row.AttributeValue("href")[1:]
			temp[i].Origin = "Netmarble"
		}
		//title
		for i, row := range titleNode {
			temp[i].Title = row.Children[0].NodeValue
		}
		for i, row := range newbieNode {
			temp[i].Newbie = row.Children[0].NodeValue
		}
		for i, row := range dateNode {
			temp[i].StartDate = row.Children[0].NodeValue[:8]
		}

		crawledData = append(crawledData, temp...)

		var currPage string
		chromedp.Sleep(2 * time.Second)
		clickerr := chromedp.Run(ctx,
			chromedp.Text("#contents > div > div > div > div.recruit_list_wrapper > div.recruit_pagination > span.page_current", &currPage),
		)
		if clickerr != nil {
			log.Fatal(clickerr)
		}

		if currPage == totalPage {
			break
		} else {
			err = chromedp.Run(ctx,
				chromedp.Click("#contents > div > div > div > div.recruit_list_wrapper > div.recruit_pagination > button.page_next", chromedp.NodeVisible),
			)

		}

		//fmt.Println(i)
	}
	for _, dat := range crawledData {
		log.Printf("%s", dat.Title)
		log.Printf("%s", dat.URL)
	}

	return crawledData
}

func BodyText(box crawler.Job) { //현재 쓸데없는 값까지 하는 중->예외처리 실패로 인해..
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
	var candNodes []*cdp.Node
	var sectionName string
	err = chromedp.Run(ctx,
		chromedp.Nodes("body > div.main > div.position-show > div > div > div.content-body.col-item.col-xs-12.col-sm-12.col-md-12.col-lg-8>section",
			&nodes),
	)
	crawler.ErrHandler(err)
	for _, row := range candNodes {
		next_html := row.FullXPath() + "/td[2]"
		var temp string
		err = chromedp.Run(ctx,
			chromedp.Text(next_html,
				&temp),
		)
		if temp == "기간" {
			err = chromedp.Run(ctx,
				chromedp.Text(row.FullXPath()+"/td[3]",
					&box.StartDate),
			)
		}

		crawler.ErrHandler(err)
	}
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
	_ = ioutil.WriteFile("./dataset/test/"+crawler.Exceptspecial(box.URL)+".json", doc, 0644)

}
