package main

import (
	"log"
	"os"

	"github.com/Vadim-Makhnev/TCP-Chat/cmd/server"
	"github.com/Vadim-Makhnev/TCP-Chat/internal/logger"
)

func main() {
	logger := logger.NewLogger("logrus")
	server := server.NewServer(logger)

	if err := server.Run(); err != nil {
		log.Fatalf("server error: %s", err)
		os.Exit(1)
	}
}
