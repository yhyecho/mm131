package service

import (
	"fmt"
	"mm131_server/api/form"
	"mm131_server/conf"
	"mm131_server/crawler"
	"mm131_server/model"
	"strconv"
	"strings"
)

// FindPicListByCategoryIDAndPageNum : 获取
func FindPicListByCategoryIDAndPageNum(categoryForm *form.CategoryForm) (ImgArr []model.ImgInfo, err error) {
	strArr := strings.Split(categoryForm.CategoryID, "_")
	var suffURL string
	if categoryForm.PageNum > 1 {
		suffURL = fmt.Sprintf(strArr[0] + "/list_" + strArr[1] + "_" + strconv.Itoa(categoryForm.PageNum) + ".html")
	} else {
		suffURL = strArr[0]
	}

	reqURL := conf.BaseURL + suffURL
	ImgArr = crawler.GetMiniPicList(reqURL, strArr[0])
	return
}

// GetImgStreamByID : 获取图片流
func GetImgStreamByID(imgID string, imgNum string) []byte {
	reqURL := fmt.Sprintf(conf.ImgReqURL, imgID, imgNum)
	return crawler.RenderPicByURL(reqURL)
}

// GetImgDetailTotalByID : 获取详情页总数
func GetImgDetailTotalByID(categoryID string, imgID string) (total int) {
	url := fmt.Sprintf(conf.BaseURL + categoryID + "/" + imgID + ".html")
	return crawler.GetTotalNumByHomeImg(url)
}
