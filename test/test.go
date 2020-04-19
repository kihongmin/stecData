package main

import (
	"log"
	"time"
)

func main() {
	//rocketpunch.Start(0)
	now := time.Now()
	custom := now.Format("2006-01-02")
	log.Println(custom)
}
