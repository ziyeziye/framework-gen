package routers

import (
	"github.com/ziyeziye/framework/pkg/config"
	"github.com/ziyeziye/framework/pkg/session"
	"github.com/gin-gonic/gin"
	"html/template"
)

func InitRouter() (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Static("/static", config.GetApp("STATIC_PATH").String())
	r.Static("/images", config.GetApp("IMAGE_PATH").String())
	r.Static("/reader", config.GetApp("PUBLIC_PATH").String()+"/reader")

	gin.SetMode(config.RumMode)
	//注册模板方法
	r.SetFuncMap(template.FuncMap{
		"splitBy" : splitBy,
		"splitJson" : splitJson,
	})
	r.LoadHTMLGlob(config.GetApp("VIEW_PATH").String())

	//session
	r.Use(session.CookieStorage(
		"mysession",
		config.GetApp("SESSION_SCREPT").String()),
	)


	viewFm(r)
	apiv1(r)
	return
}




