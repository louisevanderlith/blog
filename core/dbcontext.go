package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Articles husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Articles: husk.NewTable(Article{}, serials.GobSerial{}),
	}
}

func Shutdown() {
	ctx.Articles.Save()
}
