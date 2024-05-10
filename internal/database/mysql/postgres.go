package mysql

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Product Product
}

func New() *Mysql {
	return &Mysql{
		Product: &product{},
	}
}

func Connect() (*gorm.DB, error) {
	host := "db" // Docker-composeの場合、各サービスは独自のネットワークで実行されるため、他のサービスを探すときはサービス名で探せる。
	user := "mysql"
	password := "mysql"
	dbname := "develop"
	port := "3306"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func Close(db *gorm.DB) {
	d, errInDefer := db.DB()
	if errInDefer != nil {
		panic(errInDefer)
	}
	_ = d.Close()
}

func DBFromCtx(ctx context.Context) *gorm.DB {
	return ctx.Value("db").(*gorm.DB)
}
