package models

import (
	"fm/pkg/function/php"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         int    `gorm:"primary_key" json:"id"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (m *Model) BeforeCreate(scope *gorm.Scope) error {
	createTime := php.Date("Y:m:d H:i:s")
	scope.SetColumn("CreateTime", createTime)
	scope.SetColumn("UpdateTime", createTime)

	return nil
}

func (m *Model) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateTime", php.Date("Y:m:d H:i:s"))

	return nil
}

func ModelSearch(maps map[string]interface{}) (q *gorm.DB, e error) {
	query := db
	page, pageOk := maps["page"]
	size, sizeOk := maps["size"]
	table, tableOk := maps["table"].(string)
	order, orderOk := maps["order"].(string)
	delete(maps, "page")
	delete(maps, "size")
	delete(maps, "table")
	delete(maps, "order")

	if tableOk {
		query = query.Table(table)
	}

	if len(maps) > 0 {
		cond, vals, err := whereBuild(maps)
		if err != nil {
			return nil, err
		}

		query = db.Where(cond, vals...)
	}

	if pageOk && sizeOk {
		query = query.Offset(page).Limit(size)
	}

	if orderOk {
		query = query.Order(order)
	}

	return query, nil
}
