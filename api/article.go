package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/blog/core"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"net/http"
)

func FetchArticle(web *http.Client, host string, k hsk.Key) (core.Article, error) {
	url := fmt.Sprintf("%s/articles/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Article{}, err
	}

	defer resp.Body.Close()

	result := core.Article{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchLatestArticles(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/articles/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := records.NewResultPage(core.Article{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
