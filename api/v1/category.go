package v1

import (
	"fmt"
	"mm131_server/api/form"
	"mm131_server/defs"
	"mm131_server/model"
	"mm131_server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCategoryList 获取分类列表
func GetCategoryList(c *gin.Context) {
	var result [6]model.Category

	xinggan := model.Category{
		CategoryID:   "xinggan_6",
		CategoryName: "性感美女",
	}

	qingchun := model.Category{
		CategoryID:   "qingchun_1",
		CategoryName: "清纯美眉",
	}

	xiaohua := model.Category{
		CategoryID:   "xiaohua_2",
		CategoryName: "美女校花",
	}

	chemo := model.Category{
		CategoryID:   "chemo_3",
		CategoryName: "性感车模",
	}

	qipao := model.Category{
		CategoryID:   "qipao_4",
		CategoryName: "旗袍美女",
	}

	mingxing := model.Category{
		CategoryID:   "mingxing_5",
		CategoryName: "明星写真",
	}

	result[0] = xinggan
	result[1] = qingchun
	result[2] = xiaohua
	result[3] = chemo
	result[4] = qipao
	result[5] = mingxing

	c.JSON(http.StatusOK, &defs.JSONResult{
		Status: 0,
		Msg:    "success",
		Data:   result,
	})
	return
}

// GetHomePicListByCategoryID 根据分类和页码获取主页列表
func GetHomePicListByCategoryID(c *gin.Context) {
	var categoryForm form.CategoryForm

	err := c.BindJSON(&categoryForm)
	if err != nil {
		fmt.Printf("err of 6666 %v", err)
		return
	}
	result, err := service.FindPicListByCategoryIDAndPageNum(&categoryForm)
	if err != nil {
		fmt.Printf("err of 777 %v", err)
		return
	}
	c.JSON(http.StatusOK, &defs.JSONResult{
		Status: 0,
		Msg:    "success",
		Data:   result,
	})
	return
}

// GetImgDetailTotalByID : 获取图片总数
func GetImgDetailTotalByID(c *gin.Context) {
	categoryID := c.Query("categoryId")
	imgID := c.Query("imgId")

	result := service.GetImgDetailTotalByID(categoryID, imgID)

	c.JSON(http.StatusOK, &defs.JSONResult{
		Status: 0,
		Msg:    "success",
		Data:   result,
	})
	return

}

// GetImgStreamByID : 获取图片流
func GetImgStreamByID(c *gin.Context) {

	imgID := c.Query("imgId")
	imgNum := c.Query("imgNum")

	imgSteam := service.GetImgStreamByID(imgID, imgNum)

	c.Data(http.StatusOK, "image/jpeg", imgSteam)
	c.Done()
	return
}
