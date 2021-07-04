package biz

import (
	"context"
	"time"
)

//领域对象 DO ?
type Article struct {
	UserID    int64
	Id        int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Like      int64
}

//业务模型跟PO长得一样，说明你的业务逻辑强耦合了数据的实现细节，这个其实是你设计的问题，
//如果你是面向领域的一个object，它可能是有很大区别的，或者说有可能需要几个po对象才能够建成一个完整的领域对象，
//这个其实还是你设计上做的不好。

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

func (au *ArticleUsecase) GetArticle(id int64) (*Article, error) {
	return au.repo.GetArticle(context.Background(), id)
}

func (au *ArticleUsecase) CreateArticle(article *Article) error {
	//TODO: User 是否有权限创建文章, 决定数据是否持久化
	au.repo.CreateArticle(context.Background(), article)
	return nil
}
