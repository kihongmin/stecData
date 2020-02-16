package nexon

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func Nexon() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var nodes []*cdp.Node
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://career.nexon.com/user/recruit/notice/noticeList`),
		//chromedp.Click("#container > ul > li:nth-child(1)", chromedp.NodeVisible),
		// 인재채용 페이지까지 들어옴
		chromedp.Location(&loc),
		//chromedp.Click("#page-container > div > div.signin-wrapper > form > div.clearfix > button",chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < 16; i++ { //일단 15페이지까지 있어서 고정 설정
		var clickSelector string
		if i == 1 {
			clickSelector = "#container > ul > li:nth-child(1)"
		} else {
			clickSelector = "#con_right > div.content > div > a.next"
		}
		err := chromedp.Run(ctx,
			chromedp.Click(clickSelector, chromedp.NodeVisible),
			chromedp.Sleep(2*time.Second),
			//url을 모두 node에 저장
			chromedp.Nodes("#con_right > div.content > table > tbody > tr > td.tleft.fc_02 > a", &nodes, chromedp.ByQueryAll),
			//다음 버튼 클릭
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i)
		for _, n := range nodes {
			fmt.Printf("https://career.nexon.com%s \n", n.AttributeValue("href"))
		}
	}

	log.Printf("Landed on %s", loc)
}
