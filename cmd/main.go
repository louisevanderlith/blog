package main

import (
	"flag"
	"github.com/louisevanderlith/blog/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/blog/core"
)

func main() {
	issuer := flag.String("issuer", "http://127.0.0.1:8080/auth/realms/mango", "OIDC Provider's URL")
	audience := flag.String("audience", "blog", "Token target 'aud'")
	flag.Parse()

	core.CreateContext()
	defer core.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8102",
		Handler:      handles.SetupRoutes(*audience, *issuer),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
