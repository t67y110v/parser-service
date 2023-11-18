package main

import (
	"flag"
	"fmt"

	"github.com/t67y110v/parser-service/internal/app/apiserver"
	"github.com/t67y110v/parser-service/internal/logging"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// @title Parser service
// @version 2.0.0
// @description Swag documentaion for parser service API

// @host localhost:4000
// @BasePath /

func main() {

	logger, err := logging.NewLogger()
	if err != nil {
		panic(err)
	}

	flag.Parse()
	config := apiserver.NewConfig()
	_, err = toml.DecodeFile(configPath, config)
	if err != nil {
		logger.Error("error while decoding config file ", err)
		panic(err)
	}

	logger.Info(fmt.Sprintf("Starting apiserver addr : %s\n", config.BindAddr))
	if err := apiserver.Start(config, logger); err != nil {
		panic(err)
	}
}
