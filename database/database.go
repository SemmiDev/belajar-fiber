package database

import (
	"LearnFiber/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// DB Database instance
var DB *sql.DB

// Database settings
var (
	host     = config.Config("DB_HOST")
	port     = config.Config("DB_PORT")
	user     = config.Config("DB_USER")
	password = config.Config("DB_PASSWORD")
	dbname   = config.Config("DB_NAME")
)

// Connect function
func Connect() error {
	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		panic(err.Error())
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	fmt.Println("Connection Opened to Database")
	return nil
}