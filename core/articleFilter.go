package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"strings"
)

type articleFilter func(obj Article) bool

func (f articleFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Article))
}

func byPublished() articleFilter {
	return func(obj Article) bool {
		return obj.Public
	}
}

func byCategory(category string) articleFilter {
	lowCat := strings.ToLower(category)
	return func(obj Article) bool {
		return obj.Public && strings.ToLower(obj.Category) == lowCat
	}
}
