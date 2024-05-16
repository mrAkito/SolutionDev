package server

import (
	"SolutionDev/internal/database/mysql"
)

type Server struct {
	Tigira *tigira
}

func New(
	rdb *mysql.Mysql,
) *Server {
	return &Server{
		Tigira: &tigira{rdb},
	}
}
