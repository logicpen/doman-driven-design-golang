package main

import (
	"context"
	"ddd-golang/domain"
	_userHandler "ddd-golang/user/delivery/http"
	_userRepo "ddd-golang/user/repository/inmemory"
	_userUseCase "ddd-golang/user/useCase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	address := "127.0.0.1:9090"

	// initialize persistent storage
	// we are using local map but it can be replaced with db connection
	mapStore := make(map[string]domain.User)

	mux := http.NewServeMux()

	// initialize user repository
	userRepo := _userRepo.Init(mapStore)
	_userHandler.NewUserHandler(mux, _userUseCase.NewUserUseCase(userRepo))

	httpServer := http.Server{
		Addr:    address,
		Handler: mux,
	}

	go func() {
		log.Printf("server starting on %s", address)
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe Error: %v", err)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("HTTP Server Shutdown Error: %v", err)
	}
	log.Println("Server shutdown successful")
}
