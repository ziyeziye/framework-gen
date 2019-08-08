package routers

import (
	"fm/controllers/fm"
	"github.com/gin-gonic/gin"
)

func viewFm(r *gin.Engine) {
	r.GET("/", fm.Index)
	r.GET("/index.html", fm.Index)
	r.GET("/category.html", fm.Category)
	r.GET("/article.html", fm.Article)
	r.GET("/links.html", fm.Links)
	r.GET("/404.html", fm.NotFound)
	r.GET("/users.html", fm.Readers)
	r.GET("/tags.html", fm.Tags)
	r.GET("/search.html", fm.Search)
	r.GET("/movie/:id.html", fm.MovieTpl)
}
