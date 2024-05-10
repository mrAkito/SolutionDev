package main

import (
	"SolutionDev/internal/database/mysql"
	"SolutionDev/internal/middleware"
	"SolutionDev/internal/router"
	"SolutionDev/internal/server"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	rdb := mysql.New()

	gdb, err := mysql.Connect()
	if err != nil {
		return
	}

	defer mysql.Close(gdb)

	middleware.SetUp(e, gdb)

	srv := server.New(rdb)

	router.New(e, srv)

	e.Logger.Fatal(e.Start(":9090"))

}
