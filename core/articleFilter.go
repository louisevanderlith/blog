package core

import (
	"strings"

	"github.com/louisevanderlith/husk"
)

type articleFilter func(obj *Article) bool

func (f articleFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(*Article))
}

func byPublished() articleFilter {
	return func(obj *Article) bool {
		return obj.Public
	}
}

func byCategory(category string) articleFilter {
	lowCat := strings.ToLower(category)
	return func(obj *Article) bool {
		return obj.Public && strings.ToLower(obj.Category) == lowCat
	}
}
