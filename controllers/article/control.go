package article

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/blog/core"
	"github.com/louisevanderlith/husk"
)

func Get(c *gin.Context) {
	results := core.GetLatestArticles(1, 10)

	c.JSON(http.StatusOK, results)
}

func View(c *gin.Context) {
	k := c.Param("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	rec, err := core.GetArticle(key)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, rec)
}

func Search(c *gin.Context) {
	page, size := droxo.GetPageData(c.Param("pagesize"))

	results := core.GetNonPublicArticles(page, size)

	c.JSON(http.StatusOK, results)
}

func Create(c *gin.Context) {
	var obj core.Article
	err := c.Bind(&obj)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	rec := obj.Create()

	if rec.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, rec.Error)
	}

	c.JSON(http.StatusOK, rec)
}

func Update(c *gin.Context) {
	log.Println(c.Get("client"))
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body := &core.Article{}
	err = c.Bind(body)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = body.Update(key)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, nil)
}

func Delete(c *gin.Context) {
	log.Println(c.MustGet("client"))
	k := c.Param("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = core.RemoveArticle(key)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, "Completed")
}
