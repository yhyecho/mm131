# 图片代理服务Golang实现

## 接口说明
### 一. mm131_获取分类列表
> http://localhost:9090/getCategoryList

> GET

```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "categoryId": "xinggan_6",
            "categoryName": "性感美女"
        },
        {
            "categoryId": "qingchun_1",
            "categoryName": "清纯美眉"
        },
        {
            "categoryId": "xiaohua_2",
            "categoryName": "美女校花"
        },
        {
            "categoryId": "chemo_3",
            "categoryName": "性感车模"
        },
        {
            "categoryId": "qipao_4",
            "categoryName": "旗袍美女"
        },
        {
            "categoryId": "mingxing",
            "categoryName": "明星写真_5"
        }
    ]
}
```

### 二. mm131_根据分类和页码获取首张缩略图列表
> http://localhost:9090/getHomePicListByCategoryId

> POST

```json
{
   "categoryId": "xinggan_6",
   "pageNum": 1
}

{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "title": "梦中女神艺轩极品护士装性感破丝袜",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5096&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5096"
        },
        {
            "title": "火辣尤物周于希长腿高跟等你品尝",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5095&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5095"
        },
        {
            "title": "热情空姐卓娅祺制服诱惑姿态撩人",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5094&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5094"
        },
        {
            "title": "情趣爆乳娘妲己摆弄风骚只为等你",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5093&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5093"
        },
        {
            "title": "金发碧女梦心月破碎丝袜绝美翘臀",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5092&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5092"
        },
        {
            "title": "翘臀人妻尹菲肉色开裆丝袜秀绝美胴体",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5091&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5091"
        },
        {
            "title": "绝美尤物心妍小公主肉丝黑丝轮番轰炸",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5090&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5090"
        },
        {
            "title": "长腿女神芝芝肉丝包臀令人欲仙欲死",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5089&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5089"
        },
        {
            "title": "绝色芭比张雨萌内衣湿身春光乍泄",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5088&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5088"
        },
        {
            "title": "美女教师周妍希课后欲望高涨强行露出",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5087&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5087"
        },
        {
            "title": "贴心秘书周于希黑丝网袜等你起床爱抚",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5086&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5086"
        },
        {
            "title": "大胆御姐尹菲水晶开档丝袜太勾魂",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5085&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5085"
        },
        {
            "title": "极品御姐尤妮丝巨乳丰臀娇喘升天",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5084&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5084"
        },
        {
            "title": "乖巧喵女仆王雨纯情趣美胸寂寞难耐",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5083&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5083"
        },
        {
            "title": "蛇蝎美人黄乐然性感和服舞弄风姿",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5082&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5082"
        },
        {
            "title": "魅力娇妻何嘉颖情趣内衣等你调教",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5081&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5081"
        },
        {
            "title": "夜店女王艺轩丝袜长腿曲线很动人",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5080&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5080"
        },
        {
            "title": "童颜巨乳奈美身材火辣夺目难以抵挡",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5079&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5079"
        },
        {
            "title": "爆乳媚娘李雅不知火舞装扮大胆露出",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5078&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5078"
        },
        {
            "title": "黑丝眼镜娘周于希酒店放纵妩媚至极",
            "imageUrl": "http://localhost:9090/GetImgStreamById?imgId=5077&imgNum=0",
            "detailUrl": "http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5077"
        }
    ]
}

```

#### 三. mm131_根据图片集id或图片序号获取图片流
> http://localhost:9090/GetImgStreamById?imgId=5096&imgNum=0

> GET

> 返回图片流

#### 四. mm131_根据图片分类获取图片集大小
> http://localhost:9090/GetImgDetailTotalById?categoryId=xinggan&imgId=5096

> GET

> 
```json
{
    "status": 0,
    "msg": "success",
    "data": 51
}
```