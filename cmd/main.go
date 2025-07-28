package main

import (
	"os"

	"github.com/Vadim-Makhnev/TCP-Chat/cmd/server"
	"github.com/Vadim-Makhnev/TCP-Chat/internal/logger"
)

func main() {
	logger := logger.NewLogger("logrus")
	server := server.NewServer(logger)

	if err := server.Run(); err != nil {
		logger.Fatalf("server error: %s", err)
		os.Exit(1)
	}
}
