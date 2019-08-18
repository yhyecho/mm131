package form

// CategoryForm 分类请求参数
type CategoryForm struct {
	CategoryID string `json:"categoryId"`
	PageNum    int    `json:"pageNum"`
}
