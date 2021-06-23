package v1

//response data Struct
type ArticleDTO struct {
	Err     error `json:"err"`
	Id      int64 `json:"ArticleID"`
	Title   string
	Content []byte
}
