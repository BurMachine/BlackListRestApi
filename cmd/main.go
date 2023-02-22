package main

import (
	"blacklistApi/internal/config"
	"blacklistApi/internal/server"
	"flag"
	"log"
)

func main() {
	cfgPath := flag.String("config", "./config.yaml", "Path to yaml configuration file")
	flag.Parse()

	// Инициализация конфигурации
	conf := config.NewConfigStruct()

	err := conf.LoadConfig(*cfgPath)
	if err != nil {
		log.Fatalln(err)
	}
	serv := server.New(*conf)
	err = serv.Run()
	if err != nil {
		log.Fatalln(err)
	}

}
