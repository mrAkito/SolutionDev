package router

import (
	"SolutionDev/internal/server"
	"github.com/labstack/echo/v4"
)

type router struct {
	e   *echo.Echo
	rdb *server.Server
}

func New(
	e *echo.Echo,
	rdb *server.Server,
) {
	r := &router{e: e, rdb: rdb}
	api := e.Group("/api")
	r.group(api)
}
