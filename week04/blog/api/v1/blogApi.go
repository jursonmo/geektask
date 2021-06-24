package v1

//response data Struct, TODO: error 转换处理
type ArticleDTO struct {
	Err     error  `json:"err"`
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content []byte `json:"content"`
}

//get artcile req
type ArticleReq struct {
	Id int64
}

type ArticleCreateReq struct {
	UserId int64
	ArticleDTO
}

type ArticleCreateResp struct {
	Err error
	Id  int64
}
