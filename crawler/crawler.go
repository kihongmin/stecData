package crawler

import (
	"log"
)

type Job struct {
	URL       string
	Title     string
	Origin    string
	StartDate string
	Newbie    bool
	Content   string
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
