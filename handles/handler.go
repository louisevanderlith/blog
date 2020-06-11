package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, secureUrl string) http.Handler {
	r := mux.NewRouter()

	view := kong.ResourceMiddleware("blog.articles.view", scrt, secureUrl, ViewArticle)
	r.HandleFunc("/articles/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	get := kong.ResourceMiddleware("blog.articles.search", scrt, secureUrl, GetArticles)
	r.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}", get).Methods(http.MethodGet)
	r.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", get).Methods(http.MethodGet)

	create := kong.ResourceMiddleware("blog.articles.create", scrt, secureUrl, CreateArticle)
	r.HandleFunc("/articles", create).Methods(http.MethodPost)

	update := kong.ResourceMiddleware("blog.articles.update", scrt, secureUrl, UpdateArticle)
	r.HandleFunc("/articles", update).Methods(http.MethodPut)

	lst, err := kong.Whitelist(http.DefaultClient, secureUrl, "blog.articles.view", scrt)

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
