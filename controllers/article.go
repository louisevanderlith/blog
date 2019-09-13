package controllers

import (
	"net/http"

	"github.com/louisevanderlith/blog/core"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Article struct {
}

func (req *Article) Get(ctx context.Requester) (int, interface{}) {
	results := core.GetNonPublicArticles(1, 10)

	return http.StatusOK, results
}

// /:key
func (req *Article) View(ctx context.Requester) (int, interface{}) {
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
func (req *Article) Search(ctx context.Requester) (int, interface{}) {
	page, size := ctx.GetPageData()
	results := core.GetNonPublicArticles(page, size)

	return http.StatusOK, results
}

// @Title Create Article
// @Description Create an Article
// @Param	body		body 	core.Article	true		"body for blog article"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *Article) Create(ctx context.Requester) (int, interface{}) {
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
func (req *Article) Update(ctx context.Requester) (int, interface{}) {
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

func (req *Article) Delete(ctx context.Requester) (int, interface{}) {
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
