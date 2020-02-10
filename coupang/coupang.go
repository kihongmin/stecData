package coupang

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func Coupang() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var nodes []*cdp.Node
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://rocketyourcareer.kr.coupang.com/%ea%b2%80%ec%83%89-%ec%a7%81%eb%ac%b4`),
		//chromedp.Click("#container > ul > li:nth-child(1)", chromedp.NodeVisible),
		// 인재채용 페이지까지 들어옴
		chromedp.Location(&loc),
		//chromedp.Click("#page-container > div > div.signin-wrapper > form > div.clearfix > button",chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < 18; i++ { //일단 고정 설정
		err := chromedp.Run(ctx,
			chromedp.Sleep(2*time.Second),
			//url을 모두 node에 저장
			chromedp.Nodes("#search-results-list > ul > li > a", &nodes, chromedp.ByQueryAll),
			chromedp.Click("#pagination-bottom > div.pagination-paging > a.next", chromedp.NodeVisible),
			//다음 버튼 클릭
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i)
		for _, n := range nodes {
			fmt.Printf("https://rocketyourcareer.kr.coupang.com%s \n", n.AttributeValue("href"))
		}
	}

	log.Printf("\nLanded on %s", loc)
}
