package core

import (
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/validation"
	"html/template"
)

//Article is a Blog Post
type Article struct {
	Title     string `hsk:"size(128)"`
	Category  string
	ImageKey  keys.TimeKey
	Content   template.HTML `hsk:"size(4096)"`
	WrittenBy string        `hsk:"size(64)"`
	Public    bool          `hsk:"default(false)"`
	Intro     string        `hsk:"size(128)"`
}

func (a Article) Valid() error {
	return validation.Struct(a)
}
