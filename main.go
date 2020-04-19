package main

import (
	"geekermeter-data/crawler/rocketpunch"
	"time"
)

func main() {
	now := time.Now()
	custom := now.AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
	input_date := custom[5:7] + "/" + custom[8:10]

	rocketpunch.Start(0, input_date)
}
