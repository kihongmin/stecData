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
	Newbie    bool     `json:"newbie"`
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
