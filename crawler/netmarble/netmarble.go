package netmarble

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type data struct {
	url   string
	title string
	level string
	group string
	date  string
}

func Netmarble() []data {
	var crawledData []data

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://m.netmarble.com/rem/www/noticelist.jsp`),
		chromedp.Location(&loc),
	)
	if err != nil {
		log.Fatal(err)
	}

	//fint max page
	var totalPage string
	clickerr := chromedp.Run(ctx,
		chromedp.Text("#pageCount", &totalPage, chromedp.ByID),
	)
	if clickerr != nil {
		log.Fatal(clickerr)
	}

	for {
		temp := make([]data, 10) // 최소단위 : 10

		chromedp.Sleep(2 * time.Second)
		//url
		var nodes []*cdp.Node
		if err := chromedp.Run(ctx, chromedp.Nodes("#contents > div > div > div > div.recruit_list_wrapper > ul > li> div.cw_jopinfo > a", &nodes)); err != nil {
			log.Fatal(err)
		}
		for i, row := range nodes {
			//temp.url = "https://programmers.co.kr/" + row.AttributeValue("href")
			temp[i].url = "https://m.netmarble.com/rem/www" + row.AttributeValue("href")[1:]
		}

		//title
		var titleNode []*cdp.Node
		if err = chromedp.Run(ctx, chromedp.Nodes("#contents > div > div > div > div.recruit_list_wrapper > ul > li > div.cw_jopinfo > a > span.cw_title", &titleNode, chromedp.ByQueryAll)); err != nil {
			log.Fatal(err)
		}
		for i, row := range titleNode {
			temp[i].title = row.Children[0].NodeValue
		}

		//level
		var levelNode []*cdp.Node
		if err = chromedp.Run(ctx, chromedp.Nodes("#contents > div > div > div > div.recruit_list_wrapper > ul > li > div.cw_jopinfo > a > span.cw_info > span.cw_type", &levelNode, chromedp.ByQueryAll)); err != nil {
			log.Fatal(err)
		}
		for i, row := range levelNode {
			temp[i].level = row.Children[0].NodeValue
		}

		//group
		var groupNode []*cdp.Node
		if err = chromedp.Run(ctx, chromedp.Nodes("#contents > div > div > div > div.recruit_list_wrapper > ul > li > div.cw_group", &groupNode, chromedp.ByQueryAll)); err != nil {
			log.Fatal(err)
		}
		for i, row := range groupNode {
			temp[i].group = row.Children[0].NodeValue
		}

		//date
		var dateNode []*cdp.Node
		if err = chromedp.Run(ctx, chromedp.Nodes("#contents > div > div > div > div.recruit_list_wrapper > ul > li > div.cw_jopinfo > a > span.cw_info > span.cw_range", &dateNode, chromedp.ByQueryAll)); err != nil {
			log.Fatal(err)
		}
		for i, row := range dateNode {
			temp[i].date = row.Children[0].NodeValue
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
		log.Printf("%s", dat.title)
		log.Printf("%s", dat.url)
		log.Printf("%s", dat.level)
		log.Printf("%s", dat.group)
		log.Printf("%s", dat.date)
	}

	return crawledData
}
