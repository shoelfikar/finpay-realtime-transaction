package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/shoelfikar/finpay-realtime-transaction/controller"
	"github.com/shoelfikar/finpay-realtime-transaction/model"
	"github.com/shoelfikar/finpay-realtime-transaction/repository"
	"github.com/shoelfikar/finpay-realtime-transaction/services"
	// "github.com/shoelfikar/finpay-realtime-transaction/middleware"
)

func SetupRoutes(DB *sql.DB) *fiber.App {

	router := fiber.New(fiber.Config{
		AppName: "Finpay - Realtime Transaction",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(model.ResponseJSON{
				Error: "Internal Server Error",
				Success: "false",
				Data: struct{}{},
			})
		},
	})

	router.Use(recover.New())

	// Changing TimeZone & TimeFormat
	router.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Auth Controller
	userRepo := repository.NewUserRepository(DB)
	userService := services.NewUserService(userRepo)
	authController := controller.NewAuthController(userService)



	nonAuth := router.Group("/api/v1")
	// api := router.Group("/api/v1", middleware.JWTMiddleware)

	nonAuth.Post("/auth/login", authController.Login)
	nonAuth.Post("/auth/register", authController.Register)

	return router
}
