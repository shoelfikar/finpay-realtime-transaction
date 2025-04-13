package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/shoelfikar/finpay-realtime-transaction/model"
	"github.com/shoelfikar/finpay-realtime-transaction/services"
)

type userController struct {
	UserService services.UserService
	validate    *validator.Validate
}

type UserController interface {
	GetUserDetail(c *fiber.Ctx) error
}

func NewUserController(user services.UserService, validate *validator.Validate) UserController {
	return &userController{
		UserService: user,
		validate: validate,
	}
}

// GetUsers godoc
// @Summary Get User Detail
// @Description Get User Detail
// @Tags Users
// @Accept  json
// @Produce  json
// @Security  BearerAuth
// @Router /user/detail [get]
func (u *userController) GetUserDetail(c *fiber.Ctx) error {
	session := c.Locals("claims")
	userLocal := session.(jwt.MapClaims)
	userId := userLocal["user_id"].(string)

	user := u.UserService.GetUserDetail(userId)

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(model.ResponseJSON{
			Success: "false",
			Message: "user not found",
		})
	}

	response := model.ResponseJSON{
		Message: "Login Success",
		Success: "true",
		Data: user,
	 }
  
	 return c.JSON(response)
}
