package template

var RouterTmpl = `package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// example for init the database:
//
//  DB, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/employees?charset=utf8&parseTime=true")
//  if err != nil {
//  	panic("failed to connect database: " + err.Error())
//  }
//  defer db.Close()

var DB *gorm.DB

func ConfigRouter(router *gin.RouterGroup) {
    {{range .}}config{{pluralize .}}Router(router)
    {{end}}
}
`
