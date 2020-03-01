package main

import (
	"geekermeter-data/crawler/programmers"
)

func main() {

	//kakao.Kakao()
	//coupang.Coupang()
	//nexon.Nexon()
	//rocketpunch.BodyText("https://www.rocketpunch.com/jobs/63916/Data-Scientist")
	//crafton.Crafton()
	//naver.Naver()
	//ncsoft.Ncsoft()
	//netmarble.Netmarble()

	test := programmers.Programmers()
	for i := 0; i < 300; i++ {
		programmers.BodyText(test[i])
	}

}
