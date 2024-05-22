package mysql

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Tigira Tigira
}

func New() *Mysql {
	return &Mysql{
		Tigira: &tigira{},
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Print("No .env file found")
	}
}

func Connect() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

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
