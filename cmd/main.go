package main

import (
	"log"

	"github.com/StanislavDimitrenco/restapi"
)

func main() {
	server := new(todo.Server)

	if err := server.Run("8080"); err != nil {
		log.Fatal("error")
	}
}