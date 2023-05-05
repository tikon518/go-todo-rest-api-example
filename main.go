package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmrashed/go-todo-rest-api-example/app"
	"github.com/jmrashed/go-todo-rest-api-example/config"
)

func main() {
	cfg := config.GetConfig()

	app := &app.App{}
	app.Initialize(cfg)

	if err := http.ListenAndServe(":8080", app.Router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	} else {
		fmt.Println("Project running ...")
	}
}
