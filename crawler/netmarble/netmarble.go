package netmarble

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func Netmarble() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var nodes []*cdp.Node
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://m.netmarble.com/rem/www/noticelist.jsp`),
		chromedp.Location(&loc),
	)
	if err != nil {
		log.Fatal(err)
	}

	for i := 2; i <= 6; i++ { //일단 고정 설정->마지막페이지는 현재 안됨. 예쁘게 할려고 나중으로 미룸
		err := chromedp.Run(ctx,
			chromedp.Sleep(2*time.Second),
			chromedp.Nodes("#contents > div > div > div > div.recruit_list_wrapper > ul > li > div.cw_jopinfo > a", &nodes, chromedp.ByQueryAll),
		)
		if err != nil {
			log.Fatal(err)
		}

		for _, n := range nodes {
			fmt.Printf("https://m.netmarble.com/rem/www%s \n", n.AttributeValue("href")[1:])
		}
		clickerr := chromedp.Run(ctx,
			chromedp.Click("#contents > div > div > div > div.recruit_list_wrapper > div.recruit_pagination > button.page_next", chromedp.NodeVisible),
		)
		if err != nil {
			log.Fatal(clickerr)
		}

		//fmt.Println(i)
	}

	log.Printf("\nLanded on %s", loc)
}
