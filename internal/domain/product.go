package domain

import (
	"database/sql"
	"time"
)

type Product struct {
	Id         string         `gorm:"column:id"`
	Name       string         `gorm:"column:name"`
	Version    int            `gorm:"column:version"`
	Prefecture sql.NullString `gorm:"column:prefecture"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
}

/*
TableNameメソッドを追加しその戻り値に文字列を指定することで、
Gormが構造体名から勝手にテーブル名を推測することがなくなる。
上記の場合「products」だと認識される。
*/
func (Product) TableName() string {
	return "product"
}
