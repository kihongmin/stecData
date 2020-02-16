package rocketpunch

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func Rocketpunch() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var nodes []*cdp.Node
	var loc string
	err := chromedp.Run(ctx,

		//chromedp.Click("#container > ul > li:nth-child(1)", chromedp.NodeVisible),
		// 인재채용 페이지까지 들어옴
		chromedp.Location(&loc),
		//chromedp.Click("#page-container > div > div.signin-wrapper > form > div.clearfix > button",chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 54; i++ { //일단 고정 설정
		err := chromedp.Run(ctx,
			chromedp.Navigate(`https://www.rocketpunch.com/jobs?page=`+strconv.Itoa(i)),
			chromedp.Sleep(2*time.Second),
			//url을 모두 node에 저장
			chromedp.Nodes("#company-list > div > div.content > div.company-jobs-detail > div > div > a.nowrap.job-title.primary.link", &nodes, chromedp.ByQueryAll),
			//다음 버튼 클릭
		)
		if err != nil {
			log.Fatal(err)
		}

		for _, n := range nodes {
			fmt.Printf("https://www.rocketpunch.com/%s \n", n.AttributeValue("href"))
		}
		//fmt.Println(i)
	}

	log.Printf("\nLanded on %s", loc)
}

func main() {
	Crawler()
}
