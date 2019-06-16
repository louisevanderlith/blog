package controllers

import (
	"encoding/json"
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

// /:key
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

// @router /non/:pagesize [get]
func (req *ArticleController) GetNonPublic() {
	page, size := req.GetPageData()
	results := core.GetNonPublicArticles(page, size)

	req.Serve(http.StatusOK, nil, results)
}

// @Title Create Article
// @Description Create an Article
// @Param	body		body 	core.Article	true		"body for blog article"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *ArticleController) Post() {
	var obj core.Article
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &obj)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec := obj.Create()

	if rec.Error != nil {
		req.Serve(http.StatusInternalServerError, rec.Error, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}

// @Title UpdateArticle
// @Description Updates a Website
// @Param	body		body 	core.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *ArticleController) Put() {
	body := &core.Article{}
	key, err := req.GetKeyedRequest(body)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	err = body.Update(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, nil)
}

func (req *ArticleController) Delete() {
	k := req.Ctx.Input.Param(":key")
	key, err := husk.ParseKey(k)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	err = core.RemoveArticle(key)

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, "Completed")
}
