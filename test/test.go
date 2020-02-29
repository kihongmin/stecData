package main

import (
	"encoding/json"
	"fmt"
	"geekermeter-data/crawler"
	"io/ioutil"
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
	/*
		var test crawler.Job
		test.Title = "for test"
		test.URL = "https://programmers.co.kr/job_positions/1663"
		test.Origin = "kihong"

		programmers.BodyText(test)
	*/
	b, err := ioutil.ReadFile("./articles.json") // articles.json 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []crawler.Job // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언

	json.Unmarshal(b, &data) // JSON 문서의 내용을 변환하여 data에 저장

	fmt.Println(data)
}
