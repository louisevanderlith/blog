package routers

import (
	"fmt"
	"strings"

	"github.com/louisevanderlith/blog/controllers"
	"github.com/louisevanderlith/mango"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func Setup(s *mango.Service, host string) {
	ctrlmap := EnableFilters(s, host)
	articleCtrl := controllers.NewArticleCtrl(ctrlmap)

	beego.Router("/v1/article", articleCtrl, "post:Post;put:Put")
	beego.Router("/v1/article/:key", articleCtrl, "get:GetByKey")
	beego.Router("/v1/article/all/:pagesize", articleCtrl, "get:Get")
	beego.Router("/v1/article/non/:pagesize", articleCtrl, "get:GetNonPublic")
}

func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Admin
	emptyMap["GET"] = roletype.Unknown
	emptyMap["PUT"] = roletype.Admin

	ctrlmap.Add("/v1/article", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "PUT", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}
