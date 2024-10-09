package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PandaPy/pginer/template/initialize"
	"github.com/PandaPy/pginer/template/initialize/config"
	"github.com/PandaPy/pginer/template/router"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	initialize.Init()

	gin.SetMode(config.AppConfig.MODE)

	r := gin.New()

	router.SetupRoutes(r)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Listen),
		Handler:        r,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	color.Yellow("Server successfully running on port: %d", config.AppConfig.Listen)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %s", err)
	}
}
