package main

import (
	"github.com/ogarridojimenez/go-swagger-gin/config"
	"github.com/ogarridojimenez/go-swagger-gin/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	//TODO: Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//TODO: Initialize Router
	router.Initialice()

}
