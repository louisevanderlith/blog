package routers

import (
	"github.com/louisevanderlith/blog/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e resins.Epoxi) {
	e.JoinBundle("/", roletype.Unknown, mix.JSON, &controllers.Public{})
	e.JoinBundle("/", roletype.Admin, mix.JSON, &controllers.Article{})

	//Article
	/*
		artlCtrl := &controllers.ArticleController{}
		artlGroup := routing.NewRouteGroup("article", mix.JSON)
		artlGroup.AddRoute("Create Article", "", "POST", roletype.Admin, artlCtrl.Post)
		artlGroup.AddRoute("Update Article", "", "PUT", roletype.Admin, artlCtrl.Put)
		artlGroup.AddRoute("Article By Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, artlCtrl.GetByKey)
		artlGroup.AddRoute("Delete Article", "/{key:[0-9]+\x60[0-9]+}", "DELETE", roletype.Admin, artlCtrl.Delete)
		artlGroup.AddRoute("All Published Articles", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Unknown, artlCtrl.Get)
		artlGroup.AddRoute("All Published Articles By Category", "/all/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}", "GET", roletype.Unknown, artlCtrl.GetByCategory)
		artlGroup.AddRoute("All Articles", "/all/non/{pagesize:[A-Z][0-9]+}", "GET", roletype.Unknown, artlCtrl.GetNonPublic)
		e.AddBundle(artlGroup)*/
}
