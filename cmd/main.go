package main

import (
	"crud-task/internal/handlers"
	"net/http"
)

const (
	ADDR = ":80"
)

func run() {
	handler := handlers.NewUserHandler()

	err := http.ListenAndServe(ADDR, handler)

	if err != nil {
		panic("Failed to start HTTP server")
	}
}

func main() {
	run()
}
