package main

import (
	"os"
	"path"

	"github.com/louisevanderlith/blog/core"
	"github.com/louisevanderlith/blog/routers"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/servicetype"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	//host := os.Getenv("HOST")
	pubPath := path.Join(keyPath, pubName)

	conf, err := droxolite.LoadConfig()

	if err != nil {
		panic(err)
	}

	// Register with router
	srv := droxolite.NewService(conf.Appname, pubPath, conf.HTTPPort, servicetype.API)

	err = srv.Register()

	if err != nil {
		panic(err)
	}

	poxy := droxolite.NewEpoxy(srv)
	routers.Setup(poxy)

	core.CreateContext()
	defer core.Shutdown()

	err = poxy.Boot()

	if err != nil {
		panic(err)
	}
}
