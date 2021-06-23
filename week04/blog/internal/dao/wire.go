package dao

import "github.com/google/wire"

var ProvideSet = wire.NewSet(NewArticleRepo, NewData)
