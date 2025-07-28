package main

import (
	"fmt"
	"log"

	"github.com/victorramos887/gofinanceiro/src/config"
	"github.com/victorramos887/gofinanceiro/src/handler"
	"github.com/victorramos887/gofinanceiro/src/router"
)

func main() {
	fmt.Println("Starting the application...")

	if err := config.Init(); err != nil {
		log.Fatal("Failed to initialize configuration:", err)
	}

	handler.InitializeHandler()

	router.InitializeRoutes()
	fmt.Println("Application started successfully. On port 8085")
}
