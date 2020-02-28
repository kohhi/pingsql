package main

import (
	"database/sql"
	"errors"
	"flag"
	"os"
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
		message    bool
	)
	flag.StringVar(&dbUser, "u", envStr("DB_USER", "root"), "Database User Name.")
	flag.StringVar(&dbPassword, "p", envStr("DB_PASS", "password"), "Database User Password.")
	flag.StringVar(&dbURL, "h", envStr("DB_HOST", "localhost"), "Database Host Name.")
	flag.IntVar(&dbPort, "c", envInt("DB_PORT", 3306), "Database Connection Port.")
	flag.StringVar(&dbName, "n", envStr("DB_NAME", "database"), "Database Name.")
	flag.BoolVar(&message, "message", false, "View Message Mode.")
	flag.Parse()
	envstrings := []string{dbUser, ":", dbPassword, "@tcp(", dbURL, ":", strconv.Itoa(dbPort), ")/", dbName}
	if err := databaseRunner(envstrings, flag.Arg(0)); err != nil {
		if message {
			println(strings.Join(envstrings, ""))
			println(err.Error())
		}
		os.Exit(1)
	}
	os.Exit(0)
}

func databaseRunner(envstrings []string, arg string) error {
	database, err := sql.Open("mysql", strings.Join(envstrings, ""))
	if err != nil {
		return errors.New("failed to open database")
	}
	defer database.Close()
	database.SetConnMaxLifetime(time.Second)
	if err = database.Ping(); err != nil {
		return errors.New("failed to ping database")
	}
	if arg == "isinited" {
		if err := database.QueryRow("SELECT 1 FROM `.pingsql-initialized` LIMIT 1"); err != nil {
			return errors.New("initialized flag not found")
		}
	}
	if arg == "discard" {
		if err := database.QueryRow("DROP TABLE `.pingsql-initialized`"); err != nil {
			return errors.New("cannot delete initialized flag, or flag already deleted")
		}
	}
	if arg == "inited" {
		if err := database.QueryRow("CREATE TABLE `.pingsql-initialized`"); err != nil {
			return errors.New("cannot create initialized flag, or flag already created")
		}
	}
	return nil
}

func envStr(envName string, defStr string) string {
	str, ret := os.LookupEnv(envName)
	if !ret {
		return defStr
	}
	return str
}

func envInt(envName string, defNum int) int {
	num, err := strconv.Atoi(envStr(envName, strconv.Itoa(defNum)))
	if err != nil {
		return defNum
	}
	return num
}