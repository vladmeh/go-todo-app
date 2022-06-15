package main

import (
	"github.com/vladmeh/go-todo-app"
	"github.com/vladmeh/go-todo-app/pkg/handler"
	"github.com/vladmeh/go-todo-app/pkg/repository"
	"github.com/vladmeh/go-todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
