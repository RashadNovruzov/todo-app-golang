package main

import (
	"log"

	"github.com/RoshkANovruzov/todo-app-golang"
	"github.com/RoshkANovruzov/todo-app-golang/pkg/handler"
)

func main() {
	server := new(todo.Server)
	handler := new(handler.Handler)
	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
