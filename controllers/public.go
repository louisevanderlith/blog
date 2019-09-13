package controllers

import (
	"net/http"

	"github.com/louisevanderlith/blog/core"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Public struct {
}

func (x *Public) Get(ctx context.Requester) (int, interface{}) {
	results := core.GetLatestArticles(1, 10)

	return http.StatusOK, results
}

// @router /:pagesize [get]
func (x *Public) Search(ctx context.Requester) (int, interface{}) {
	page, size := ctx.GetPageData()
	results := core.GetLatestArticles(page, size)

	return http.StatusOK, results
}

func (req *Public) View(ctx context.Requester) (int, interface{}) {
	k := ctx.FindParam("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec, err := core.GetArticle(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, rec
}

//this should use search !
// @router /all/:category/:pagesize [get]
func (req *Public) GetByCategory(ctx context.Requester) (int, interface{}) {
	category := ctx.FindParam("category")
	page, size := ctx.GetPageData()
	results := core.GetArticlesByCategory(category, page, size)

	return http.StatusOK, results
}
