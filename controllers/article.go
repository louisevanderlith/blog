package controllers

import (
	"net/http"

	"github.com/louisevanderlith/blog/core"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/control"
)

type ArticleController struct {
	control.APIController
}

func NewArticleCtrl(ctrlmap *control.ControllerMap) *ArticleController {
	result := &ArticleController{}
	result.SetInstanceMap(ctrlmap)

	return result
}

// /v1/artcicle/:key
func (req *ArticleController) GetByKey() {
	k := req.Ctx.Input.Param(":key")
	key, err := husk.ParseKey(k)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec, err := core.GetArticle(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}

// @router /all/:pagesize [get]
func (req *ArticleController) Get() {
	page, size := req.GetPageData()
	results := core.GetLatestArticles(page, size)

	req.Serve(http.StatusOK, nil, results)
}
