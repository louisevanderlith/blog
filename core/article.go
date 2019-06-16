package core

import (
	"github.com/louisevanderlith/husk"
)

type Article struct {
	Title     string `hsk:"size(128)"`
	ImageKey  husk.Key
	Content   string `hsk:"size(4096)"`
	WrittenBy string `hsk:"size(64)"`
	Public    bool   `hsk:"default(false)"`
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
	return ctx.Articles.Find(page, size, byPublished())
}

func GetNonPublicArticles(page, size int) husk.Collection {
	return ctx.Articles.Find(page, size, husk.Everything())
}

func RemoveArticle(key husk.Key) error {
	//defer ctx.Articles.Save()
	return ctx.Articles.Delete(key)
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

	obj.Set(a)

	defer ctx.Articles.Save()
	return ctx.Articles.Update(obj)
}

func (a Article) Publish(key husk.Key) error {
	a.Public = true

	return a.Update(key)
}
