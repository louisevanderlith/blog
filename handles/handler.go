package handles

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
)

func SetupRoutes(host, scrt, secureUrl string) http.Handler {
	r := mux.NewRouter()

	get := kong.ResourceMiddleware("blog.articles.search", scrt, secureUrl, GetArticles)
	r.HandleFunc("/articles", get).Methods(http.MethodGet)

	view := kong.ResourceMiddleware("blog.articles.view", scrt, secureUrl, ViewArticle)
	r.HandleFunc("/articles/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	srch := kong.ResourceMiddleware("blog.articles.search", scrt, secureUrl, SearchArticles)
	r.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}", srch).Methods(http.MethodGet)
	r.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	create := kong.ResourceMiddleware("blog.articles.create", scrt, secureUrl, CreateArticle)
	r.HandleFunc("/articles", create).Methods(http.MethodPost)

	update := kong.ResourceMiddleware("blog.articles.update", scrt, secureUrl, UpdateArticle)
	r.HandleFunc("/articles", update).Methods(http.MethodPut)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //[]string{fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))}, //you service is available and allowed for this base url
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
