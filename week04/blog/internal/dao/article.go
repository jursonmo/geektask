package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jursonmo/geektask/week04/blog/internal/biz"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

var articleID int

func (ar *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	articleID++
	//TODO:
	return nil
}

func (ar *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	//TODO : ar.data get data from dataDriver

	a := &biz.Article{
		Id:        id,
		Title:     fmt.Sprintf("Title%d", id),
		Content:   fmt.Sprintf("Content:%d", id),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return a, nil
}
