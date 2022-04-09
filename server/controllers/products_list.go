package controller

import (
	"github.com/gin-gonic/gin"
	service "github.com/lil-shimon/oms/services"
)

/// Index function for products_list
func Index(c *gin.Context) {
	ps, err := service.GetProductsLists()
	if err != nil {
		c.AbortWithStatus(404)
		panic(err)
	}

	c.JSON(200, ps)
}
