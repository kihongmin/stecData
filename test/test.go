package main

import (
	"geekermeter-data/crawler"
	"geekermeter-data/crawler/nexon"
)

func main() {

	//kakao.Kakao()
	//coupang.Coupang()
	var box crawler.Job
	box.URL = "https://career.nexon.com/user/recruit/notice/noticeView?joinCorp=NX&reNo=20170363"
	nexon.BodyText(box)
	//rocketpunch.Start()
	//crafton.Crafton()
	//naver.Naver()
	//ncsoft.Ncsoft()
	//netmarble.Netmarble()

}
