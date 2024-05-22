package server

import (
	"SolutionDev/internal/database/mysql"
)

type Server struct {
	Tigira *tigira
	Video *video
}

func New(
	rdb *mysql.Mysql,
) *Server {
	return &Server{
		Tigira: &tigira{rdb},
		Video: &video{rdb},
	}
}
