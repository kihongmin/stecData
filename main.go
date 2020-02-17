package main

import (
	"geekermeter-data/crawler/programmers"
	"geekermeter-data/db"
)

func main() {
	//kakao.Kakao()
	//coupang.Coupang()
	//nexon.Nexon()
	//rocketpunch.Rocketpunch()
	//crafton.Crafton()
	//naver.Naver()
	//ncsoft.Ncsoft()
	// netmarble.Netmarble()
	programmersJobs := programmers.Programmers()
	db.Insert(programmersJobs)
}
