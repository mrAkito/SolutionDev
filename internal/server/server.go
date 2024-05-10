package server

import (
	"SolutionDev/internal/database/mysql"
)

type Server struct {
	Product *product
}

func New(
	rdb *mysql.Mysql,
) *Server {
	return &Server{
		Product: &product{rdb},
	}
}
