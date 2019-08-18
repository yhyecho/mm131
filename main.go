package main

import (
	v1 "mm131_server/api/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getCategoryList", v1.GetCategoryList)
	router.POST("/getHomePicListByCategoryId", v1.GetHomePicListByCategoryID)
	router.GET("/GetImgStreamById", v1.GetImgStreamByID)
	router.GET("/GetImgDetailTotalById", v1.GetImgDetailTotalByID)
	router.Run(":9090")
}
