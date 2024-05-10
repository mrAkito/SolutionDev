package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *router) group(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello echo")
	})

	product := g.Group("/product")
	{
		product.GET("", r.rdb.Product.List)
		product.GET("/:id", r.rdb.Product.Get)
	}
}
