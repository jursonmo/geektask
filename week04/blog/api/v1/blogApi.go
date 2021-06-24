package v1

//response data Struct, TODO: error 转换处理
type ArticleDTO struct {
	Err     error  `json:"err"`
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content []byte `json:"content"`
}

//
type ArticleReq struct {
	Id int64
}
