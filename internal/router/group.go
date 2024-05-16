package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *router) group(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello echo")
	})

	tigira := g.Group("/tigira")
	{
		tigira.POST("/:goal_name", r.rdb.Tigira.Create)
	}
}
