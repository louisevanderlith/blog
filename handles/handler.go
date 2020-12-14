package handles

import (
	"github.com/louisevanderlith/droxolite/open"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupRoutes(audience, issuer string) http.Handler {
	r := mux.NewRouter()
	mw := open.BearerMiddleware(audience, issuer)
	r.Handle("/articles", mw.Handler(http.HandlerFunc(GetArticles))).Methods(http.MethodGet)
	r.Handle("/articles/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewArticle))).Methods(http.MethodGet)

	r.Handle("/articles/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchArticles))).Methods(http.MethodGet)
	r.Handle("/articles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchArticles))).Methods(http.MethodGet)

	r.Handle("/articles", mw.Handler(http.HandlerFunc(CreateArticle))).Methods(http.MethodPost)
	r.Handle("/articles/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateArticle))).Methods(http.MethodPut)

	//lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "blog.articles.view", scrt)

	//if err != nil {
	//	panic(err)
	//}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
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
