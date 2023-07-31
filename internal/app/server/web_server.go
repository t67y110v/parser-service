package server

import (
	"github.com/t67y110v/parser-service/internal/app/config"
	"github.com/t67y110v/parser-service/internal/app/logging"
)

func Start(config *config.Config) error {

	logger := logging.GetLogger()

	server := newServer(config, logger)

	//StartServerWithGracefulShutdown(server, config.BindAddr)
	return server.router.Listen(config.Port)
}
