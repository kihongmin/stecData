package programmers

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type data struct {
	url         string
	title       string
	companyname string
}

func Programmers() []data {
	var crawledData []data
	//crawledData := make([]data, 200)
	//var temp data
	count := 0

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://programmers.co.kr/job`),
		chromedp.Location(&loc),
	)
	if err != nil {
		log.Fatal(err)
	}
	for {
		temp := make([]data, 20) // 최소단위 : 20

		chromedp.Sleep(2 * time.Second)
		//url & title node
		var nodes []*cdp.Node
		if err := chromedp.Run(ctx, chromedp.Nodes("#list-positions-wrapper > ul > li > div.item-body > h4 > a", &nodes)); err != nil {
			log.Fatal(err)
		}
		for i, row := range nodes {
			//temp.url = "https://programmers.co.kr/" + row.AttributeValue("href")
			temp[i].url = "https://programmers.co.kr/" + row.AttributeValue("href")
			temp[i].title = row.Children[0].NodeValue
		}

		chromedp.Sleep(2 * time.Second)

		//company-name node
		var nameNode []*cdp.Node
		if err = chromedp.Run(ctx, chromedp.Nodes("#list-positions-wrapper > ul > li> div.item-body > h5", &nameNode, chromedp.ByQueryAll)); err != nil {
			log.Fatal(err)
		}
		for i, row := range nameNode {
			count++
			temp[i].companyname = row.Children[0].NodeValue
		}

		crawledData = append(crawledData, temp...)
		// move to next
		var checkor []*cdp.Node

		chromedp.Sleep(2 * time.Second)
		clickerr := chromedp.Run(ctx,
			//chromedp.Click("#paginate > nav > ul > li.next.next_page.page-item > a", chromedp.NodeVisible),
			chromedp.Nodes("#paginate > nav > ul > li.next.next_page.page-item > a", &checkor),
		)
		if clickerr != nil {
			log.Fatal(clickerr)
		}
		if checkor[0].AttributeValue("rel") != "next" {
			break
		} else {
			err = chromedp.Run(ctx,
				chromedp.Click("#paginate > nav > ul > li.next.next_page.page-item > a", chromedp.NodeVisible),
			)
		}

	}

	for _, dat := range crawledData {
		log.Printf("%s", dat.title)
		log.Printf("%s", dat.url)
		log.Printf("%s", dat.companyname)
	}

	log.Println(count, cap(crawledData), len(crawledData))
	return crawledData
}
