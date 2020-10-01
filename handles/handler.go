package handles

import (
	"github.com/louisevanderlith/kong/middle"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupRoutes(scrt, securityUrl, managerUrl string) http.Handler {
	r := mux.NewRouter()
	ins := middle.NewResourceInspector(http.DefaultClient, securityUrl, managerUrl)
	get := ins.Middleware("blog.articles.search", scrt, GetArticles)
	r.HandleFunc("/articles", get).Methods(http.MethodGet)

	view := ins.Middleware("blog.articles.view", scrt, ViewArticle)
	r.HandleFunc("/articles/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	srch := ins.Middleware("blog.articles.search", scrt, SearchArticles)
	r.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}", srch).Methods(http.MethodGet)
	r.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	create := ins.Middleware("blog.articles.create", scrt, CreateArticle)
	r.HandleFunc("/articles", create).Methods(http.MethodPost)

	update := ins.Middleware("blog.articles.update", scrt, UpdateArticle)
	r.HandleFunc("/articles", update).Methods(http.MethodPut)

	lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "blog.articles.view", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
