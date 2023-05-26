package main

import (
	"log"
	"task/pkg/config"
	"task/pkg/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	config, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("failed to load config env", config)
	}

	r := gin.Default()
	handler.Routes(r)
	err := r.Run(config.Port)
	if err != nil {
		log.Fatal("failed to start server", config)
	}
}
