package controllers

import (
	"net/http"

	"github.com/louisevanderlith/blog/core"
	"github.com/louisevanderlith/droxolite/context"
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

//this should use search !
// @router /all/:category/:pagesize [get]
func (req *Public) GetByCategory(ctx context.Requester) (int, interface{}) {
	category := ctx.FindParam("category")
	page, size := ctx.GetPageData()
	results := core.GetArticlesByCategory(category, page, size)

	return http.StatusOK, results
}
