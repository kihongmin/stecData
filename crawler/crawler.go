package crawler

import (
	"log"
	"regexp"
	//"geekermeter-data/crawler/naver"
	//"github.com/caarlos0/env"
)

type Job struct {
	URL       string   `json:"url"`
	Title     string   `json:"title"`
	Origin    string   `json:"origin"`
	StartDate string   `json:"start_date"`
	Newbie    []int8   `json:"newbie"` // 신입, 경력 둘다의 경우
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
type Newbie int8

const (
	intern = 1 + iota
	newbie
	intnew
	expert
	intexp
	newexp
	intnewexp
)

var newbies = [...]string{
	"intern",    // 001
	"newbie",    // 010
	"intnew",    // 011
	"expert",    // 100
	"intexp",    // 101
	"newexp",    // 110
	"intnewexp", //111
}

func (n Newbie) String() string { return newbies[(n-1)%7] }

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
	re := regexp.MustCompile(`[가-힣]+|\s|/`)
	key := re.ReplaceAllString(word, "")
	return key
}

func Getnewbie(word string) []int8 {
	//인턴 10 신입 50 경력 90
	level := make([]int8, 0, 4)
	i := regexp.MustCompile(`인턴`)
	n := regexp.MustCompile(`신입`)
	e := regexp.MustCompile(`경력`)
	dm := regexp.MustCompile(`경력 무관`)
	if dm.MatchString(word) {
		return append(level, 50, 90)
	}
	if i.MatchString(word) {
		level = append(level, 10)
	}
	if n.MatchString(word) {
		level = append(level, 50)
	}
	if e.MatchString(word) {
		level = append(level, 90)
	}

	return level
}
