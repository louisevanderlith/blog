package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
)

type BlogContext interface {
	UpdateArticle(k hsk.Key, obj Article) error
	CreateArticle(obj Article) (hsk.Key, error)
	RemoveArticle(key hsk.Key) error

	GetArticle(key hsk.Key) (Article, error)
	GetLatestArticles(page, size int) (records.Page, error)
	GetNonPublicArticles(page, size int) (records.Page, error)
}

type context struct {
	Articles husk.Table
}

func (c context) UpdateArticle(k hsk.Key, obj Article) error {
	return c.Articles.Update(k, obj)
}

func (c context) CreateArticle(obj Article) (hsk.Key, error) {
	return c.Articles.Create(obj)
}

func (c context) GetArticle(key hsk.Key) (Article, error) {
	rec, err := ctx.Articles.FindByKey(key)

	if err != nil {
		return Article{}, err
	}

	return rec.GetValue().(Article), nil
}

func (c context) GetLatestArticles(page, size int) (records.Page, error) {
	return c.Articles.Find(page, size, byPublished())
}

func (c context) GetNonPublicArticles(page, size int) (records.Page, error) {
	return c.Articles.Find(page, size, op.Everything())
}

func (c context) RemoveArticle(key hsk.Key) error {
	return c.Articles.Delete(key)
}

var ctx context

func Context() BlogContext {
	return ctx
}

func CreateContext() {
	ctx = context{
		Articles: husk.NewTable(Article{}),
	}
}

func Shutdown() {
	ctx.Articles.Save()
}
