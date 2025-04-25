package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/atm_simulation"
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Database connected successfully")
}
