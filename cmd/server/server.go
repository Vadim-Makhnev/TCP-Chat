package server

import (
	"fmt"
	"net"

	"github.com/Vadim-Makhnev/TCP-Chat/internal/logger"
)

type IServer interface {
	Run() error
}

type Server struct {
	logger logger.ILogger
}

func NewServer(logger logger.ILogger) IServer {
	return &Server{
		logger,
	}
}

func (s *Server) Run() error {
	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		return fmt.Errorf("server started with error: %w", err)
	}

	s.logger.Info("server started on", conn.Addr())

	for {
		_, err := conn.Accept()
		if err != nil {
			s.logger.Error(err.Error())
			continue
		}
	}
}
