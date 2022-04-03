package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	DBMigrate(DBConnect())
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "成功",
		})
	})
	r.Run(":1323")
}

type User struct {
	gorm.Model
	Name string `json:"name"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}

func DBConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "oms"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	return db
}
