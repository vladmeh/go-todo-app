package main

import (
	"github.com/spf13/viper"
	"github.com/vladmeh/go-todo-app"
	"github.com/vladmeh/go-todo-app/pkg/handler"
	"github.com/vladmeh/go-todo-app/pkg/repository"
	"github.com/vladmeh/go-todo-app/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
