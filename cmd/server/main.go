package main

import (
	apiserver2 "LongTaskAPI/internal/apiserver"
	"LongTaskAPI/internal/repository/in_memory"
	"LongTaskAPI/internal/services"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	pathToConfig string
)

func main() {
	flag.StringVar(&pathToConfig, "path", "config/apiserver.toml", "Path to config file")
	flag.Parse()

	cfg := apiserver2.NewConfig()

	_, err := toml.DecodeFile(pathToConfig, cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := in_memory.NewInMemoryTaskRepo()
	service := services.NewTaskService(repo)

	api := apiserver2.New(cfg, service)

	err = api.Run()
	if err != nil {
		log.Fatal(err)
	}
}
