package ncsoft

// 클릭 후 내용 크롤링까지 하고 중단. 다른거부터
import (
	"context"
	"geekermeter-data/crawler"
	"log"

	"github.com/chromedp/chromedp"
)

var (
	baseURL = "https://careers.ncsoft.com/apply/list"
)

func Ncsoft() []crawler.Job {
	var crawledData []crawler.Job

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(baseURL),
		chromedp.Location(&loc),
	)
	crawler.ErrHandler(err)

	clickerr := chromedp.Run(ctx,
		chromedp.Click("#container > section > article > ul > li:nth-child(1) > a.panel.slat-bttn.applyDetailBtn", chromedp.NodeVisible),
	)
	crawler.ErrHandler(clickerr)

	var res string
	var loc1 string
	err = chromedp.Run(ctx,
		chromedp.Location(&loc1),
		chromedp.Text("#container > section > article > section:nth-child(3) > div > p", &res, chromedp.NodeVisible),
	)
	crawler.ErrHandler(err)
	log.Printf("\n%s", res)
	log.Printf("\n%s", loc1)

	/*
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
	*/
	//url node
	return crawledData
}
