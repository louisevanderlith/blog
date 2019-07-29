package routers

import (
	"github.com/louisevanderlith/blog/controllers"
	"github.com/louisevanderlith/droxolite"

	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(poxy *droxolite.Epoxy) {
	//Article
	artlCtrl := &controllers.ArticleController{}
	artlGroup := droxolite.NewRouteGroup("article", artlCtrl)
	artlGroup.AddRoute("/", "POST", roletype.Admin, artlCtrl.Post)
	artlGroup.AddRoute("/", "PUT", roletype.Admin, artlCtrl.Put)
	artlGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, artlCtrl.GetByKey)
	artlGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "DELETE", roletype.Admin, artlCtrl.Delete)
	artlGroup.AddRoute("/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Unknown, artlCtrl.Get)
	artlGroup.AddRoute("/all/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}", "GET", roletype.Unknown, artlCtrl.GetByCategory)
	artlGroup.AddRoute("/all/non/{pagesize:[A-Z][0-9]+}", "GET", roletype.Unknown, artlCtrl.GetNonPublic)
	poxy.AddGroup(artlGroup)
	/*ctrlmap := EnableFilters(s, host)
	articleCtrl := controllers.NewArticleCtrl(ctrlmap)

	beego.Router("/v1/article", articleCtrl, "post:Post;put:Put")
	beego.Router("/v1/article/:key", articleCtrl, "get:GetByKey;delete:Delete")
	beego.Router("/v1/article/all/:pagesize", articleCtrl, "get:Get")
	beego.Router("/v1/article/all/:category/:pagesize", articleCtrl, "get:GetByCategory")
	beego.Router("/v1/article/non/:pagesize", articleCtrl, "get:GetNonPublic")*/
}

/*
func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Admin
	emptyMap["GET"] = roletype.Unknown
	emptyMap["PUT"] = roletype.Admin
	emptyMap["DELETE"] = roletype.Admin

	ctrlmap.Add("/v1/article", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
	}), false)

	return ctrlmap
}
*/
