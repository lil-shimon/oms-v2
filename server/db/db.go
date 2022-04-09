package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConnectする関数
// sqlboiler
// *sql.DBをreturnする
func DBConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/oms?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	return db
}

var (
	DB *gorm.DB
)

func Connect() *gorm.DB {
	// DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "oms"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"

	// db, err := gorm.Open(DBMS, CONNECT)
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetDB() *gorm.DB {
	return DB
}
