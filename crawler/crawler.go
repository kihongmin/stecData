package crawler

import (
	"log"
	"regexp"
	//"github.com/caarlos0/env"
)

type Job struct {
	URL       string   `json:"url"`
	Title     string   `json:"title"`
	Origin    string   `json:"origin"`
	StartDate string   `json:"start_date"`
	Newbie    string   `json:"newbie"` // 신입, 경력 둘다의 경우
	Content   []string `json:"content"`
}

type URLs struct {
	ID     int
	Title  string
	Origin string
	// start_date string
	// end_date string
	// position string
	URL string
	// basic string
	// advanced string
}

// errHandler is errHandler
func ErrHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ExtractNum(word string) string {
	re := regexp.MustCompile(`[^0-9]+`)
	key := re.ReplaceAllString(word, "")
	return key
}

func Exceptspecial(word string) string {
	re := regexp.MustCompile(`[^0-9A-Za-z]+`)
	key := re.ReplaceAllString(word, "")
	return key
}

func OnlyKorean(word string) string {
	re := regexp.MustCompile(`[^가-힣]+|[만원]+|[최대]+`)
	key := re.ReplaceAllString(word, "")
	return key
}

func ExceptKorean(word string) string {
	re := regexp.MustCompile(`[가-힣]+`)
	key := re.ReplaceAllString(word, "")
	return key
}
