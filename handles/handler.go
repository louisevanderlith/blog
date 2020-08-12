package handles

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
)

func SetupRoutes(scrt, securityUrl, managerUrl string) http.Handler {
	r := mux.NewRouter()

	get := kong.ResourceMiddleware(http.DefaultClient, "blog.articles.search", scrt, securityUrl, managerUrl, GetArticles)
	r.HandleFunc("/articles", get).Methods(http.MethodGet)

	view := kong.ResourceMiddleware(http.DefaultClient, "blog.articles.view", scrt, securityUrl, managerUrl, ViewArticle)
	r.HandleFunc("/articles/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	srch := kong.ResourceMiddleware(http.DefaultClient, "blog.articles.search", scrt, securityUrl, managerUrl, SearchArticles)
	r.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}", srch).Methods(http.MethodGet)
	r.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	create := kong.ResourceMiddleware(http.DefaultClient, "blog.articles.create", scrt, securityUrl, managerUrl, CreateArticle)
	r.HandleFunc("/articles", create).Methods(http.MethodPost)

	update := kong.ResourceMiddleware(http.DefaultClient, "blog.articles.update", scrt, securityUrl, managerUrl, UpdateArticle)
	r.HandleFunc("/articles", update).Methods(http.MethodPut)

	lst, err := kong.Whitelist(http.DefaultClient, securityUrl, "blog.articles.view", scrt)

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
