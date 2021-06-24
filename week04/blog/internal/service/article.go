package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/jursonmo/geektask/week04/blog/api/v1"
	"github.com/jursonmo/geektask/week04/blog/internal/biz"
)

type ArticleService struct {
	log *log.Helper
	au  *biz.ArticleUsecase
}

func NewArticleService(logger log.Logger, au *biz.ArticleUsecase) *ArticleService {
	return &ArticleService{
		log: log.NewHelper(logger),
		au:  au,
	}
}

func (as *ArticleService) GetArticle(id int) *v1.ArticleDTO {
	as.log.Infof("GetArticle")
	article, err := as.au.GetArticle(id)
	return &v1.ArticleDTO{
		Err:     err,
		Id:      article.Id,
		Title:   article.Title,
		Content: []byte(article.Content),
	}
}

func (as *ArticleService) CreateArticle() *v1.ArticleDTO {
	as.log.Infof("CreateArticle")
	err := as.au.CreateArticle()
	return &v1.ArticleDTO{
		Err: err,
	}
}
