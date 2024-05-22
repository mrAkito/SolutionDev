package server

import (
	"SolutionDev/internal/database/mysql"
	"SolutionDev/internal/server/request"
	"github.com/labstack/echo/v4"
	"net/http"
	"fmt"
)

type Video interface {
	GetVideo(c echo.Context) error
}

type video struct {
	rdb *mysql.Mysql
}

func (s *video) GetVideo(c echo.Context) error {
	var req request.PostSteps
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	ctx := c.Request().Context()

	steps := req.ToSteps()
	stepIds := make([]int, len(steps))
	for i, step := range steps {
		fmt.Print(step)
		resStep, err := s.rdb.Video.GetStep(ctx, step)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Database error")
		}
		stepIds[i] = resStep.Id
	}

	video, err := s.rdb.Video.GetVideo(ctx, request.ToVideo(stepIds))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Database error")
	}

	return c.JSON(http.StatusOK, video)
}