package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/lil-shimon/oms/controllers"
	"github.com/lil-shimon/oms/db"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	// DB connect
	db.Connect()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "成功",
		})
	})

	r.GET("/products_lists", controller.Index)
	r.Run(":1323")
}
