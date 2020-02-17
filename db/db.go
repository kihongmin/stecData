package db

import (
	"database/sql"
	"geekermeter-data/crawler"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB is DB
func connect() *sql.DB {
	configuration := crawler.GetConfig()
	mySQL := configuration.MySQLUser + ":" + configuration.MySQLPW + "@tcp(" + configuration.MySQLHost + ":" + configuration.MySQLPort + ")/" + configuration.MySQLDB
	db, err := sql.Open("mysql", mySQL)
	crawler.ErrHandler(err)
	return db
}

func Insert(jobs []crawler.Job) {
	db := connect()
	defer db.Close()

	satement, err := db.Prepare("INSERT INTO stec.Urls (title, origin, url) VALUES (?, ?, ?)")
	crawler.ErrHandler(err)

	var checker int64 = 0
	var meter int64 = 0
	for _, job := range jobs {
		ret, err := satement.Exec(job.Title, job.Origin, job.URL)
		crawler.ErrHandler(err)
		n, err := ret.RowsAffected()
		checker = checker + n
		meter++
	}
	if meter != checker {
		log.Printf("number of insert rows are not exact: %d != %d", meter, checker)
	} else {
		log.Printf("number of insert rows: %d", checker)
	}
}
