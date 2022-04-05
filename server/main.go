package main

import (
	"github.com/gin-gonic/gin"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// db.DBConnect()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "成功",
		})
	})
	// r.GET("/users", handler.GetUsers)
	r.Run(":1323")
}
