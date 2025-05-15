package main

import (
	"log"
	"net/http"

	handler "github.com/Infamia2334/go-tasks/internal/delivery/http/handlers"
	route "github.com/Infamia2334/go-tasks/internal/delivery/http/routes"
	"github.com/Infamia2334/go-tasks/internal/services"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	taskService := services.NewTaskService()

	taskHandler := handler.NewTaskHandler(taskService)

	route.RegisterRoutes(router, taskHandler)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}
