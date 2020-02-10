package wemakeprice

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func Wemakeprice() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var nodes []*cdp.Node
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://recruit.wemakeprice.com/notice/list#`),
		chromedp.Location(&loc),
	)
	if err != nil {
		log.Fatal(err)
	}


		//fmt.Println(i)
	}

	log.Printf("\nLanded on %s", loc)
}
