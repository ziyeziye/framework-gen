package template

var ControllerTmpl = `package api

import (
	"{{.Package}}/pkg/request"
	"{{.Package}}/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"

	"{{.PackageName}}"
)

func config{{pluralize .StructName}}Router(router *gin.RouterGroup) {
	router.GET("/{{pluralize .StructName | toLower}}", GetAll{{pluralize .StructName}})
	router.POST("/{{pluralize .StructName | toLower}}", Add{{.StructName}})
	router.GET("/{{pluralize .StructName | toLower}}/:id", Get{{.StructName}})
	router.PUT("/{{pluralize .StructName | toLower}}/:id", Update{{.StructName}})
	router.DELETE("/{{pluralize .StructName | toLower}}/:id", Delete{{.StructName}})
}

func GetAll{{pluralize .StructName}}(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var total int
	total = models.Get{{.StructName}}Total(maps)
	data["total"] = total

	maps = request.GetPage(c, maps, false)

	respJson := response.Json()
	if {{pluralize .StructName | toLower}}, err := models.Get{{pluralize .StructName}}(maps); err != nil {
		respJson.SetState(false).SetMsg("error")
	} else {
		data["list"] = {{pluralize .StructName | toLower}}
		respJson.SetData(data)
	}

	respJson.Return(c)
}

func Get{{.StructName}}(c *gin.Context) {
	id := request.GetParam(c, "id", "")
	{{.StructName | toLower}}, _ := models.Get{{.StructName}}(id.MustInt())

	respJson := response.Json()
	if {{.StructName | toLower}}.ID > 0 {
		respJson.SetData({{.StructName | toLower}})
	} else {
		respJson.SetState(false).SetCode(http.StatusNotFound)
	}
	respJson.Return(c)
}

func Add{{.StructName}}(c *gin.Context) {
	//maps := make(map[string]interface{})

	{{.StructName | toLower}} := models.{{.StructName}}{}
	json := response.JsonType{}

	respJson := response.Json()
	if err := models.Add{{.StructName}}(&{{.StructName | toLower}}); err != nil {
		respJson.Set(http.StatusInternalServerError, "新增失败", false, {{.StructName | toLower}})
	} else {
		respJson.SetData({{.StructName | toLower}})
	}
	respJson.Return(c)
}

func Update{{.StructName}}(c *gin.Context) {
	id := request.GetParam(c, "id", "")
	maps := make(map[string]interface{})

	{{.StructName | toLower}}, _ := models.Get{{.StructName}}(id.MustInt())

	respJson := response.Json()
	if {{.StructName | toLower}}.ID > 0 {
		if err := models.Update{{.StructName}}(&{{.StructName | toLower}}, maps); err != nil {
			respJson.Set(http.StatusInternalServerError, "修改失败", false, {{.StructName | toLower}})
		} else {
			respJson.SetData({{.StructName | toLower}})
		}
	} else {
		respJson.SetState(false).SetCode(http.StatusNotFound)
	}
	respJson.Return(c)
}

func Delete{{.StructName}}(c *gin.Context) {
	id := request.GetParam(c, "id", "")
	respJson := response.Json()
	if err := models.Delete{{.StructName}}(id.MustInt()); err != nil {
		respJson.Set(http.StatusInternalServerError, "删除失败", false, nil)
	}
	respJson.Return(c)
}
`
