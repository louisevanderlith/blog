package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Articles husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Articles: husk.NewTable(Article{}),
	}
}

func Shutdown() {
	ctx.Articles.Save()
}
