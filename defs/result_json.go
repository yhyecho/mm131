package defs

// JSONResult 统一返回对象
type JSONResult struct {
	Status int8        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
