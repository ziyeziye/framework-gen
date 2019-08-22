package template

var ModelTmpl = `package {{.PackageName}}

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"

	"github.com/guregu/null"
)

var (
    _ = time.Second
    _ = sql.LevelDefault
    _ = null.Bool{}
)


type {{.StructName}} struct {
    {{range .Fields}}{{.}}
    {{end}}
}

// TableName sets the insert table name for this struct type
func ({{.ShortStructName}} *{{.StructName}}) TableName() string {
	return tablePrefix + "{{.TableName}}"
}

func (*{{.StructName}}) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateTime", null.TimeFrom(time.Now().Local()))
	scope.SetColumn("UpdateTime", null.TimeFrom(time.Now().Local()))
	return nil
}

func (*{{.StructName}}) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateTime", null.TimeFrom(time.Now().Local()))
	return nil
}

func Get{{.StructName}}(id int) ({{.StructName | toLower}} {{.StructName}}, err error) {
	err = db.First(&{{.StructName | toLower}}, id).Error
	return {{.StructName | toLower}}, err
}

func Get{{pluralize .StructName}}(maps map[string]interface{}) ({{pluralize .StructName | toLower}} []{{.StructName}}, err error) {
	q, e := ModelSearch(maps)
	if e!=nil {
		return
	}

	err = q.Find(&{{pluralize .StructName | toLower}}).Error
	return
}

func Get{{.StructName}}Total(maps map[string]interface{}) (count int) {
	cond, vals, err := whereBuild(maps)
	if err != nil {
		return 0
	}
	db.Model(&{{.StructName}}{}).Where(cond, vals...).Count(&count)

	return
}

func Delete{{.StructName}}(id int) (err error) {
	err = db.Where("id = ?", id).Delete(&{{.StructName}}{}).Error
	return
}

func Update{{.StructName}}({{.StructName | toLower}} *{{.StructName}}, maps map[string]interface{}) (err error) {
	err = db.Model({{.StructName | toLower}}).Updates(maps).Error
	return
}


func Add{{.StructName}}({{.StructName | toLower}} *{{.StructName}}) (err error) {
	err = db.Create({{.StructName | toLower}}).Error
	return
}

`
