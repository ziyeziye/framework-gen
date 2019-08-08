package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ziyeziye/framework-gen/api"
)

func apiv1(r *gin.Engine) {
	apiv1 := r.Group("")
	{
		api.ConfigRouter(apiv1)
	}

	reader(r)
}

func reader(r *gin.Engine)  {
}
