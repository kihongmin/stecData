package main

import (
	"geekermeter-data/crawler"
	"geekermeter-data/crawler/rocketpunch"
	"time"
)

func main() {
	//rocketpunch.Start(0)
	now := time.Now()
	datetime := crawler.Exceptspecial(now.Format("2006-01-02"))
	rocketpunch.Start(0, datetime)
}
