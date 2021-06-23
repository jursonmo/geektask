package biz

import (
	"context"
	"time"
)

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

type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

func (au *ArticleUsecase) GetArticle() (*Article, error) {

	return nil, nil
}

func (au *ArticleUsecase) CreateArticle() error {
	return nil
}
