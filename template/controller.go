package template

var ControllerTmpl = `package api

import (
	"github.com/gin-gonic/gin"
	"{{.Package}}/pkg/request"
	"{{.Package}}/pkg/response"
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
	data["list"] = models.Get{{pluralize .StructName}}(maps)

	response.Json(c, response.JsonType{Data: data})
}

func Get{{.StructName}}(c *gin.Context) {
	id := request.GetParam(c, "id", "")
	{{.StructName | toLower}}, _ := models.Get{{.StructName}}(id.MustInt())

	if {{.StructName | toLower}}.ID > 0 {
		response.Json(c, response.JsonType{Data: {{.StructName | toLower}}})
	} else {
		response.Json(c, response.JsonType{
			State: false,
			Code:  http.StatusNotFound,
		})
	}
}

func Add{{.StructName}}(c *gin.Context) {
	//maps := make(map[string]interface{})

	{{.StructName | toLower}} := models.{{.StructName}}{}
	json := response.JsonType{}

	if err := models.Add{{.StructName}}(&{{.StructName | toLower}}); err != nil {
		json.Code = http.StatusInternalServerError
		json.Msg = "新增失败"
		json.State = false
		json.Data = {{.StructName | toLower}}
	}
	response.Json(c, json)
}

func Update{{.StructName}}(c *gin.Context) {
	id := request.GetParam(c, "id", "")
	maps := make(map[string]interface{})

	{{.StructName | toLower}}, _ := models.Get{{.StructName}}(id.MustInt())

	json := response.JsonType{}
	if {{.StructName | toLower}}.ID > 0 {
		if err := models.Update{{.StructName}}(&{{.StructName | toLower}}, maps); err != nil {
			json.Code = http.StatusInternalServerError
			json.Msg = "修改失败"
			json.State = false
		}
	} else {
		json.Code = http.StatusNotFound
		json.State = false
	}

	response.Json(c, json)
}

func Delete{{.StructName}}(c *gin.Context) {
	id := request.GetParam(c, "id", "")
	json := response.JsonType{}
	if err := models.Delete{{.StructName}}(id.MustInt()); err != nil {
		json.Code = http.StatusInternalServerError
		json.Msg = "删除失败"
		json.State = false
	}
	response.Json(c, json)
}
`
