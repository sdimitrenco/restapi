package main

import (
	"github.com/StanislavDimitrenco/restapi/pkg/handler"
	"github.com/StanislavDimitrenco/restapi/pkg/repository"
	"github.com/StanislavDimitrenco/restapi/pkg/service"
	"log"

	"github.com/StanislavDimitrenco/restapi"
)

func main() {

	repository := repository.NewRepository()
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)

	if err := server.Run("8080", handlers.InitRouts()); err != nil {
		log.Fatal("error")
	}
}
