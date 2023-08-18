package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	initApp "github.com/bryanck29/be-test/init"
	"github.com/bryanck29/be-test/internal/config"

	_ "github.com/bryanck29/be-test/docs"
)

// @title			API Docs
// @version		1.0
// @description	Swagger API Docs.
func main() {
	server := initApp.InitApp()

	go func() {
		serverPort := ":" + strconv.Itoa(config.Core.ServerPort)
		log.Print("Starting server...")
		server.Logger.Fatal(server.Start(serverPort))
		log.Print("Server is running")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Print("Shutting down server...")
		log.Fatal(err)
	}
}
