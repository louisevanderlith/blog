package core

import (
	"github.com/louisevanderlith/husk"
)

type Article struct {
	Title     string `hsk:"size(128)"`
	ImageKey  husk.Key
	Content   string `hsk:"size(4096)"`
	WrittenBy string `hsk:"size(64)"`
}

func (a Article) Valid() (bool, error) {
	return husk.ValidateStruct(&a)
}

func GetArticle(key husk.Key) (*Article, error) {
	rec, err := ctx.Articles.FindByKey(key)

	if err != nil {
		return nil, err
	}

	return rec.Data().(*Article), nil
}

func GetLatestArticles(page, size int) husk.Collection {
	return ctx.Articles.Find(page, size, husk.Everything())
}

func (a Article) Create() husk.CreateSet {
	return ctx.Articles.Create(a)
}

func (a Article) Update(key husk.Key) error {
	obj, err := ctx.Articles.FindByKey(key)

	if err != nil {
		return err
	}

	obj.Set(a)

	defer ctx.Articles.Save()
	return ctx.Articles.Update(obj)
}
