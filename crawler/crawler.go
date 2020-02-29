package crawler

import (
	"log"
	"regexp"

	"github.com/caarlos0/env"
)

type Job struct {
	URL       string
	Title     string
	Origin    string
	StartDate string
	Newbie    bool
	Content   string
}
type BodyText struct {
	Position     string
	Requirements string
	Preference   string
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
