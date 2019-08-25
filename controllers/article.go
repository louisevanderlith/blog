package controllers

import (
	"net/http"

	"github.com/louisevanderlith/blog/core"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type ArticleController struct {
}

// /:key
func (req *ArticleController) GetByKey(ctx context.Contexer) (int, interface{}) {
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

// @router /all/:pagesize [get]
func (req *ArticleController) Get(ctx context.Contexer) (int, interface{}) {
	page, size := ctx.GetPageData()
	results := core.GetLatestArticles(page, size)

	return http.StatusOK, results
}

// @router /non/:pagesize [get]
func (req *ArticleController) GetNonPublic(ctx context.Contexer) (int, interface{}) {
	page, size := ctx.GetPageData()
	results := core.GetNonPublicArticles(page, size)

	return http.StatusOK, results
}

// @router /all/:category/:pagesize [get]
func (req *ArticleController) GetByCategory(ctx context.Contexer) (int, interface{}) {
	category := ctx.FindParam("category")
	page, size := ctx.GetPageData()
	results := core.GetArticlesByCategory(category, page, size)

	return http.StatusOK, results
}

// @Title Create Article
// @Description Create an Article
// @Param	body		body 	core.Article	true		"body for blog article"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *ArticleController) Post(ctx context.Contexer) (int, interface{}) {
	var obj core.Article
	err := ctx.Body(&obj)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec := obj.Create()

	if rec.Error != nil {
		return http.StatusInternalServerError, rec.Error
	}

	return http.StatusOK, rec
}

// @Title UpdateArticle
// @Description Updates a Website
// @Param	body		body 	core.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *ArticleController) Put(ctx context.Contexer) (int, interface{}) {
	body := &core.Article{}
	key, err := ctx.GetKeyedRequest(body)

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = body.Update(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}

func (req *ArticleController) Delete(ctx context.Contexer) (int, interface{}) {
	k := ctx.FindParam("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = core.RemoveArticle(key)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, "Completed"
}
