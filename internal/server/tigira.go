package server

import (
	"SolutionDev/internal/database/mysql"
	"SolutionDev/internal/server/request"
	"SolutionDev/internal/server/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Tigira interface {
	Create(c echo.Context) error
}

type tigira struct {
	rdb *mysql.Mysql
}

func (s *tigira) Create(c echo.Context) error {
	var req request.PostGPSInfo
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	ctx := c.Request().Context()
	goalId, err := s.rdb.Tigira.GetGoalId(ctx, req.ToGPSInfoGoalId())
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Database error")
	}

	err = s.rdb.Tigira.Update(ctx, req.ToSearchGPSInfo(goalId.GoalId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Database error")
	}

	err = s.rdb.Tigira.Create(ctx, req.ToGPSInfos(goalId.GoalId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Database error")
	}

	return c.JSON(http.StatusOK, "Success")
}

func (s *tigira) Get(c echo.Context) error {
	var req request.GetGPSInfo
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	ctx := c.Request().Context()

	goalId, err := s.rdb.Tigira.GetGoalId(ctx, req.ToGPSInfo())
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Database error")
	}

	res, err := s.rdb.Tigira.List(ctx, req.ToGoal(goalId.GoalId, req.Option))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Database error")
	}
	return c.JSON(http.StatusOK, response.ToGoals(res))
}
