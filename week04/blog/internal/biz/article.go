package biz

import (
	"context"
	"time"
)

//领域对象
type Article struct {
	Id        int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Like      int64
}

type ArticleRepo interface {
	CreateArticle(ctx context.Context, article *Article) error
	GetArticle(ctx context.Context, id int64) (*Article, error)
}

//领域服务,实现部分业务逻辑，决定数据是否持久化, 贫血模式?
type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

func (au *ArticleUsecase) GetArticle(id int) (*Article, error) {

	return nil, nil
}

func (au *ArticleUsecase) CreateArticle() error {
	//TODO: 决定数据是否持久化
	return nil
}
