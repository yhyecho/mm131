package crawler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mm131_server/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

// ConvertToString : 字符编码转换
func ConvertToString(src string, srcCode string, tagCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

// GetMiniPicList 获取分类列表页缩略图列表
func GetMiniPicList(reqURL string, category string) (miniPicList []model.ImgInfo) {
	res, err := http.Get(reqURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("err url : %s", reqURL)
	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	htmlDom, _ := ioutil.ReadAll(res.Body)
	result := ConvertToString(string(htmlDom), "gbk", "utf-8")
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(result))
	if err != nil {
		log.Fatal(err)
		return
	}

	miniPicList = make([]model.ImgInfo, 0)
	doc.Find(".list-left dd").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		imgDom := s.Find("img")
		imgURL, _ := imgDom.Attr("src")
		title, _ := imgDom.Attr("alt")
		if imgURL == "" || title == "" {
			return
		}

		imgID := strings.Split(imgURL, "/")[4]
		imgURL = fmt.Sprintf("http://localhost:9090/GetImgStreamById?imgId=%s&imgNum=0", imgID)
		detailURL := fmt.Sprintf("http://localhost:9090/GetImgDetailTotalById?categoryId=%s&imgId=%s", category, imgID)
		miniPic := model.ImgInfo{
			Title:     title,
			ImageURL:  imgURL,
			DetailURL: detailURL,
		}

		miniPicList = append(miniPicList, miniPic)
	})

	return miniPicList
}

// GetTotalNumByHomeImg : 获取图片详情页总数
func GetTotalNumByHomeImg(url string) (total int) {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	htmlDom, _ := ioutil.ReadAll(res.Body)
	result := ConvertToString(string(htmlDom), "gbk", "utf-8")
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(result))
	if err != nil {
		log.Fatal(err)
	}
	numStr := doc.Find(".page-ch").First().Text()

	// baseURL, _ := doc.Find(".content-pic img").Attr("src")

	// baseURL = baseURL[0 : strings.LastIndex(baseURL, ".jpg")-1]
	k, _ := strconv.Atoi(numStr[3:5])

	return k
}

// RenderPicByURL : 获取图片流
func RenderPicByURL(reqURL string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("Referer", `http://www.mm131.com/`)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("Request Img Stream Err Of Code %d\n", res.StatusCode)
		return nil
	}
	imgData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Parse Img Err Of %v", err)
		return nil
	}

	return imgData
}
