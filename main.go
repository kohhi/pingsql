package main

import (
	"database/sql"
	"flag"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		dbUser     string
		dbPassword string
		dbURL      string
		dbPort     int
		dbName     string
	)
	flag.StringVar(&dbUser, "u", "root", "Database User Name.")
	flag.StringVar(&dbPassword, "p", "root", "Database User Password.")
	flag.StringVar(&dbURL, "h", "localhost", "Database Host Name.")
	flag.IntVar(&dbPort, "c", 3306, "Database Connection Port.")
	flag.StringVar(&dbName, "n", "mysql", "Database Name.")
	flag.Parse()
	envstrings := []string{dbUser, ":", dbPassword, "@tcp(", dbURL, ":", strconv.Itoa(dbPort), ")/", dbName}
	healthDatabase, err := sql.Open("mysql", strings.Join(envstrings, ""))
	if err != nil {
		panic("Failed to Connect Database.")
	}
	defer healthDatabase.Close()
	healthDatabase.SetConnMaxLifetime(time.Second)
	if err = healthDatabase.Ping(); err != nil {
		panic("Failed to Send Ping.")
	}
	return
}
