package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/blog/controllers/article"
	"github.com/louisevanderlith/blog/core"
	"github.com/louisevanderlith/blog/droxo"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	r := gin.Default()
	//r.Use(cors)
	//host := os.Getenv("HOST")
	//authority := "https://oauth2." + host
	//provider, err := oidc.NewProvider(context.Background(), authority)
	//if err != nil {
	//	panic(err)
	//}

	r.GET("/article/:key", article.View)

	authed := r.Group("/article")
	authed.Use(droxo.Authorize())
	authed.POST("", article.Create)
	authed.PUT("/:key", article.Update)
	authed.DELETE("/:key", article.Delete)

	r.GET("/articles", article.Get)
	r.GET("/articles/:pagesize/*hash", article.Search)
	err := r.Run(":8102")

	if err != nil {
		panic(err)
	}
}

/*
func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	host := os.Getenv("HOST")
	httpport, _ := strconv.Atoi(os.Getenv("HTTPPORT"))
	appName := os.Getenv("APPNAME")
	pubPath := path.Join(keyPath, pubName)

	// Register with router
	srv := bodies.NewService(appName, "", pubPath, host, httpport, servicetype.API)

	routr, err := do.GetServiceURL("", "Router.API", false)

	if err != nil {
		panic(err)
	}

	err = srv.Register(routr)

	if err != nil {
		panic(err)
	}

	poxy := resins.NewMonoEpoxy(srv, element.GetNoTheme(host, srv.ID, "none"))
	routers.Setup(poxy)
	poxy.EnableCORS(host)

	core.CreateContext()
	defer core.Shutdown()

	err = droxolite.Boot(poxy)

	if err != nil {
		panic(err)
	}
}
*/
