package main

import (
	"github.com/t67y110v/parser-service/internal/app/config"
	"github.com/t67y110v/parser-service/internal/app/logging"
	"github.com/t67y110v/parser-service/internal/app/server"
)

// @title parser-service
// @version 1.0

// @host localhost:4000
// @BasePath /

func main() {

	logging.Init()
	l := logging.GetLogger()
	l.Infoln("Config initialization")
	config, err := config.LoadConfig()
	if err != nil {
		l.Fatal(err)
	}
	l.Infof("Starting apiserver addr : %s\n", config.Port)
	if err := server.Start(&config); err != nil {
		l.Fatal(err)
	}

}
