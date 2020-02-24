package main

import (
	"geekermeter-data/crawler/programmers"
	"log"
)

func main() {
	//kakao.Kakao()
	//coupang.Coupang()
	//nexon.Nexon()
	//rocketpunch.Rocketpunch()
	//crafton.Crafton()
	//naver.Naver()
	//ncsoft.Ncsoft()
	//netmarble.Netmarble()

	k := programmers.BodyText("https://programmers.co.kr/job_positions/1610")
	for _, row := range k {
		log.Println(row.Position)
		log.Println(row.Preference)
		log.Println(row.Requirements)
	}
}
