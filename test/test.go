package main

import (
	"geekermeter-data/crawler/netmarble"
	"geekermeter-data/crawler/nexon"
	"geekermeter-data/crawler/rocketpunch"
)

func main() {

	//kakao.Kakao()
	//coupang.Coupang()
	//var box crawler.Job
	//box.URL = "https://m.netmarble.com/rem/www/notice.jsp?anno_id=1583010&annotype=all"
	//netmarble.Start()
	//nexon.Start()
	//rocketpunch.Start()
	//crafton.Crafton()
	//naver.Naver()
	//ncsoft.Ncsoft()
	//netmarble.Netmarble()
	//programmers.Start()

	forname := netmarble.Start(0)
	forname = nexon.Start(forname)
	forname = rocketpunch.Start(forname)
}
