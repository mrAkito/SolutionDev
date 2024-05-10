package server

import (
	"SolutionDev/internal/database/mysql"
	"SolutionDev/internal/server/request"
	"SolutionDev/internal/server/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Product interface {
	List(c echo.Context) error
	Get(c echo.Context) error
}

type product struct {
	rdb *mysql.Mysql
}

func (s *product) List(c echo.Context) error {
	ctx := c.Request().Context()
	products, err := s.rdb.Product.List(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.ToProducts(products))
}

func (s *product) Get(c echo.Context) error {
	var req request.GetProduct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	ctx := c.Request().Context()
	getProduct, err := s.rdb.Product.Get(ctx, req.ToProduct())
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Database error")
	}
	return c.JSON(http.StatusOK, response.ToProduct(getProduct))
}
