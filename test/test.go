package main

import (
	"geekermeter-data/crawler"
	"geekermeter-data/crawler/netmarble"
)

func main() {

	//kakao.Kakao()
	//coupang.Coupang()
	var box crawler.Job
	box.URL = "https://m.netmarble.com/rem/www/notice.jsp?anno_id=1583010&annotype=all"
	netmarble.BodyText(box)
	//rocketpunch.Start()
	//crafton.Crafton()
	//naver.Naver()
	//ncsoft.Ncsoft()
	//netmarble.Netmarble()
	//programmers.Start()

}
