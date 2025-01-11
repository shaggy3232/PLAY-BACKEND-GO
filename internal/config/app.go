package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Connect() {
	dsn := "user:user_password@tcp(mysql_container:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *sql.DB {

	return db
}
