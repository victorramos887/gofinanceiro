// @title           API Financeiro
// @version         1.0
// @description     API para gerenciamento financeiro pessoal
// @termsOfService  http://example.com/terms/

// @contact.name   Victor Ramos
// @contact.email  victor@email.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8085
// @BasePath  /api/v1

package main

import (
	"fmt"
	"log"

	"github.com/victorramos887/gofinanceiro/src/config"
	_ "github.com/victorramos887/gofinanceiro/src/docs"
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
