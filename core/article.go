package core

import (
	"html/template"

	"github.com/louisevanderlith/husk"
)

//Article is a Blog Post
type Article struct {
	Title     string `hsk:"size(128)"`
	Category  string
	ImageKey  husk.Key
	Content   template.HTML `hsk:"size(4096)"`
	WrittenBy string        `hsk:"size(64)"`
	Public    bool          `hsk:"default(false)"`
	Intro     string        `hsk:"size(128)"`
}

func (a Article) Valid() error {
	return husk.ValidateStruct(&a)
}

func GetArticle(key husk.Key) (Article, error) {
	rec, err := ctx.Articles.FindByKey(key)

	if err != nil {
		return Article{}, err
	}

	return rec.Data().(Article), nil
}

func GetLatestArticles(page, size int) (husk.Collection, error) {
	return ctx.Articles.Find(page, size, byPublished())
}

func GetNonPublicArticles(page, size int) (husk.Collection, error) {
	return ctx.Articles.Find(page, size, husk.Everything())
}

func GetArticlesByCategory(category string, page, size int) (husk.Collection, error) {
	return ctx.Articles.Find(page, size, byCategory(category))
}

func RemoveArticle(key husk.Key) error {
	err := ctx.Articles.Delete(key)

	if err != nil {
		return err
	}

	return ctx.Articles.Save()
}

func (a Article) Create() husk.CreateSet {
	a.Public = false

	defer ctx.Articles.Save()
	return ctx.Articles.Create(a)
}

func (a Article) Update(key husk.Key) error {
	obj, err := ctx.Articles.FindByKey(key)

	if err != nil {
		return err
	}

	err = obj.Set(a)

	if err != nil {
		return err
	}

	err = ctx.Articles.Update(obj)

	if err != nil {
		return err
	}

	return ctx.Articles.Save()
}
