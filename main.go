package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()

	log.Println(time.Parse("2006-01-02 15:04:05", now))

	//netmarble.Start(0, input_date)
	//db.Insert(programmersJobs)
}
