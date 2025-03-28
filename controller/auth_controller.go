package controller

import (

	"github.com/gofiber/fiber/v2"
	"github.com/shoelfikar/finpay-realtime-transaction/model"
)

// GetUsers godoc
// @Summary Login
// @Description Login user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Router /auth/login [get]
func Login(c *fiber.Ctx) error {
	response := model.ResponseJSON{
      Message: "Login Success",
      Success: "true",
      Data: struct{}{},
   }

   return c.JSON(response)
}

// GetUsers godoc
// @Summary Register
// @Description Register new user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Router /auth/register [get]
func Register(c *fiber.Ctx) error {
   response := model.ResponseJSON{
      Message: "Register User Success",
      Success: "true",
      Data: struct{}{},
   }

   return c.JSON(response)
}