package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/blog/core"
	"github.com/louisevanderlith/husk"
)

func GetArticles(w http.ResponseWriter, r *http.Request) {
	results, err := core.GetNonPublicArticles(1, 10)

	if err != nil {
		log.Println("Get Articles Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// /:key
func ViewArticle(w http.ResponseWriter, r *http.Request) {
	k := drx.FindParam(r, "key")
	key, err := husk.ParseKey(k)

	if err != nil {
		log.Println("Parse Key Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.GetArticle(key)

	if err != nil {
		log.Println("Get Article Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(rec))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @router /all/:pagesize [get]
func SearchArticles(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)
	results, err := core.GetNonPublicArticles(page, size)

	if err != nil {
		log.Println("Get Articles Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title Create Article
// @Description Create an Article
// @Param	body		body 	core.Article	true		"body for blog article"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	obj := core.Article{}
	err := drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := obj.Create()

	if err != nil {
		log.Println("Create Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(rec))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title UpdateArticle
// @Description Updates a Website
// @Param	body		body 	core.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	key, err := husk.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println("Parse Key Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := &core.Article{}
	err = drx.JSONBody(r, body)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = body.Update(key)

	if err != nil {
		log.Println("Update Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(nil))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	k := drx.FindParam(r, "key")
	key, err := husk.ParseKey(k)

	if err != nil {
		log.Println("Parse Key Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.RemoveArticle(key)

	if err != nil {
		log.Println("Remove Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON("Completed"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
