package main

import (
	"blacklistApi/internal/config"
	"blacklistApi/internal/database"
	"blacklistApi/internal/server"
	"flag"
	"github.com/rs/zerolog"
	"os"
)

// @title BlackListApi App API
// @version 1.0
// @description API Server for BlackListApi Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cfgPath := flag.String("config", "./config.yaml", "Path to yaml configuration file")
	flag.Parse()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	// Инициализация конфигурации
	conf := config.NewConfigStruct()

	err := conf.LoadConfig(*cfgPath)
	if err != nil {
		logger.Fatal().Err(err)
	}

	storage, err := database.InitConn(*conf)
	if err != nil {
		logger.Fatal().Err(err)
	}

	serv := server.New(*conf, storage)
	serv.Handlers.Logger = &logger
	err = serv.Run()
	if err != nil {
		logger.Fatal().Err(err)
	}

}
