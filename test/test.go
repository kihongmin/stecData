package main

import (
<<<<<<< HEAD
=======
	"geekermeter-data/crawler"
>>>>>>> develop
	"geekermeter-data/crawler/netmarble"
)

func main() {

	//kakao.Kakao()
	//coupang.Coupang()
<<<<<<< HEAD
	//var box crawler.Job
	//box.URL = "https://m.netmarble.com/rem/www/notice.jsp?anno_id=1583010&annotype=all"
	netmarble.Start()
	//nexon.Start()
=======
	var box crawler.Job
	box.URL = "https://m.netmarble.com/rem/www/notice.jsp?anno_id=1583010&annotype=all"
	netmarble.BodyText(box)
>>>>>>> develop
	//rocketpunch.Start()
	//crafton.Crafton()
	//naver.Naver()
	//ncsoft.Ncsoft()
	//netmarble.Netmarble()
	//programmers.Start()

}
