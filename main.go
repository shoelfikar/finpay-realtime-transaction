package main

import (
	"log"

	"github.com/gofiber/swagger"
	"github.com/shoelfikar/finpay-realtime-transaction/config"
	"github.com/shoelfikar/finpay-realtime-transaction/router"

   _ "github.com/shoelfikar/finpay-realtime-transaction/docs"
)

// @title           Swagger Finpay Realtime Transaction API
// @version         1.0
// @description     Finpay Realtime Transaction API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1

func main() {
   DB := config.InitDB()
   
   router := routes.SetupRoutes(DB)

   router.Get("/swagger/*", swagger.HandlerDefault)

   err := router.Listen(":8000")

   if err != nil {
      log.Fatalf("error running project %v", err)
   }
}