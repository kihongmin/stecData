package naver

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
	date  string
}

func Naver() []data{//아직 개발 직군만 크롤링임.
	crawledData := make([]data, 200)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://recruit.navercorp.com/naver/job/list/developer`),
		chromedp.Location(&loc),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nLanded on %s", loc)

	for {
		chromedp.Sleep(2 * time.Second)
		clickerr := chromedp.Run(ctx,
			chromedp.Click("#moreDiv > button", chromedp.NodeVisible),
		)
		if clickerr != nil {
			break
		}
	}
	chromedp.Sleep(2 * time.Second)
	log.Printf("\nclick success")

	//url node
	var nodes []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes("#jobListDiv > ul > li > a", &nodes)); err != nil {
		log.Fatal(err)
	}
	for i, k := range nodes {
		crawledData[i].url = "https://recruit.navercorp.com" + k.AttributeValue("href")
	}

	chromedp.Sleep(2 * time.Second)

	//title node
	var titleNodes []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes("#jobListDiv > ul > li > a > span > strong", &titleNodes, chromedp.ByQueryAll)); err != nil {
		log.Fatal(err)
	}
	for i, row := range titleNodes {
		for _, c := range row.Children {
			crawledData[i].title = c.NodeValue
		}
	}

	chromedp.Sleep(2 * time.Second)

	//date node
	var dateNode []*cdp.Node
	if err = chromedp.Run(ctx, chromedp.Nodes("#jobListDiv > ul > li > a > span > em", &dateNode, chromedp.ByQueryAll)); err != nil {
		log.Fatal(err)
	}
	for i, row := range dateNode {
		for _, c := range row.Children {
			crawledData[i].date = c.NodeValue
		}
	}

	for _, dat := range crawledData {
		log.Printf("%s", dat.title)
		log.Printf("%s", dat.url)
		log.Printf("%s", dat.date)
	}
	return crawledData
}
