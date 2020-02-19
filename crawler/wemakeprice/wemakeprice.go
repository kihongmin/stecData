package wemakeprice

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func Wemakeprice() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://recruit.wemakeprice.com/notice/list#`),
		chromedp.Location(&loc),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nLanded on %s", loc)
}
