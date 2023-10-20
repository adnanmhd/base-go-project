package main

import (
	"base-go-project/controller"
	"base-go-project/util"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		if panicMessage := recover(); panicMessage != nil {
			fmt.Printf("Panic '%v' captured ", panicMessage)
		}
	}()
	appConfig := util.InitConfig()
	log.Info("Configuring Log...")
	// init log file
	util.ConfigLog(appConfig.AppName)
	log.Info("===========================================")

	if strings.ToLower(appConfig.Debug) == "true" {
		log.Info("Log level set to DEBUG level")
		log.SetLevel(log.DebugLevel)
	}
	log.Infof("Starting %s apps on port %s", appConfig.AppName, appConfig.Port)
	log.Info("Configuring API router...")
	router := controller.SetupRouter(appConfig)
	router.HideBanner = true

	// Start server
	serverPort := fmt.Sprintf(":%s", appConfig.Port)
	router.Server.Addr = serverPort

	go func() {
		if err := router.Start(serverPort); err != nil {
			router.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := router.Shutdown(ctx); err != nil {
		router.Logger.Fatal(err)
	}

}
