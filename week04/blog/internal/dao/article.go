package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jursonmo/geektask/week04/blog/internal/biz"
)

//PO for ORM, 跟数据库的表结构关联?
type ArticleData struct {
	Articler string //作者

	Id        int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Like      int64
}

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

func (ar *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	// get data from dataDriver
	var articleData *ArticleData
	articleData = ar.data.GetArticleById(id)

	//deep copy: PO --> DO
	a := &biz.Article{
		Id:        articleData.Id,
		Title:     fmt.Sprintf("Title%d", articleData.Id),
		Content:   fmt.Sprintf("Content:%d", articleData.Id),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return a, nil
}

func (ar *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	articleID++
	//TODO: DO-->PO
	return nil
}
