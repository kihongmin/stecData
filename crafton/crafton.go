package crafton

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func Crafton() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var nodes []*cdp.Node
	var loc string
	err := chromedp.Run(ctx,

		// 인재채용 페이지까지 들어옴
		chromedp.Navigate(`https://krafton.recruiter.co.kr/app/jobnotice/list`),
		chromedp.Location(&loc),
	)
	if err != nil {
		log.Fatal(err)
	}

	for i := 2; i <= 6; i++ { //일단 고정 설정->마지막페이지는 현재 안됨. 예쁘게 할려고 나중으로 미룸
		err := chromedp.Run(ctx,
			chromedp.Sleep(2*time.Second),
			chromedp.Nodes("#divJobnoticeList > ul > li > div > h2 > a", &nodes, chromedp.ByQueryAll),
		)
		if err != nil {
			log.Fatal(err)
		}

		for _, n := range nodes {
			fmt.Printf("https://krafton.recruiter.co.kr/app/jobnotice/view?systemKindCode=MRS2&jobnoticeSn=%s \n", n.AttributeValue("data-jobnoticesn"))
		}
		clickerr := chromedp.Run(ctx,
			chromedp.Click("#content > div.paging-wrapper.middle-set > div > ul > li:nth-child("+strconv.Itoa(i)+") > a", chromedp.NodeVisible),
		)
		if err != nil {
			log.Fatal(clickerr)
		}

		//fmt.Println(i)
	}

	log.Printf("\nLanded on %s", loc)
}
