package service

import (
	"time"

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

func (as *ArticleService) GetArticle(id int64) *v1.ArticleDTO {
	as.log.Infof("GetArticle")
	article, err := as.au.GetArticle(id)
	_ = err
	return &v1.ArticleDTO{
		//Err:     err,
		Id:      article.Id,
		Title:   article.Title,
		Content: article.Content,
	}
}

func (as *ArticleService) CreateArticle(acr *v1.ArticleCreateReq) *v1.ArticleCreateResp {
	as.log.Infof("CreateArticle")
	//deep copy: DTO-->DO
	ba := &biz.Article{
		UserID:    acr.UserId,
		Id:        acr.Id,
		Title:     acr.Title,
		Content:   string(acr.Content),
		CreatedAt: time.Now(),
	}
	err := as.au.CreateArticle(ba)
	_ = err
	return &v1.ArticleCreateResp{
		//Err: err,
		Id: acr.Id,
	}
}
