package controller

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/shoelfikar/finpay-realtime-transaction/model"
	"github.com/shoelfikar/finpay-realtime-transaction/services"
	"github.com/shoelfikar/finpay-realtime-transaction/utils"
)

type authController struct {
   UserService services.UserService
   validate    *validator.Validate
}

type AuthController interface {
   Login(c *fiber.Ctx) error
   Register(c *fiber.Ctx) error
}

func NewAuthController(user services.UserService, validate *validator.Validate) AuthController {
   return &authController{
      UserService: user,
      validate: validate,
   }
}

// GetUsers godoc
// @Summary Login
// @Description Login user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body model.LoginRequest true "Login"
// @Router /auth/login [post]
func (a *authController) Login(c *fiber.Ctx) error {
   body := c.Body()
   var request model.LoginRequest
   err := a.validate.Struct(request)
   if ok := utils.ValidationErrors(err); ok {
      message := utils.GetErrorMessagevalidator(err)
      return c.Status(fiber.StatusBadRequest).JSON(model.ResponseJSON{
         Success: "false",
         Message: "validation error",
         ValidationError: message,
      })
   }

   err = json.Unmarshal(body, &request)
   if err != nil {
      return err
   }

   user := a.UserService.GetUserByEmail(request.Email)
   if user == nil {
      return c.Status(fiber.StatusNotFound).JSON(model.ResponseJSON{
         Message: "User not found",
         Success: "false",
         Data: struct{}{},
      })
   }

   checkPassword := utils.CheckPassword(*user.Password, request.Password)
   if !checkPassword {
      return c.Status(fiber.StatusBadRequest).JSON(model.ResponseJSON{
         Message: "Wrong password, try again",
         Success: "false",
         Data: struct{}{},
      })
   }
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
// @Param request body model.RegisterRequest true "Register"
// @Router /auth/register [post]
func (a *authController) Register(c *fiber.Ctx) error {
   body := c.Body()
   var (
      request model.RegisterRequest
   )
   err := json.Unmarshal(body, &request)
   if err != nil {
      return err
   }

   if request.Password != request.RetypePassword {
      return c.Status(fiber.StatusBadRequest).JSON(model.ResponseJSON{
         Message: "Password confirmation not match",
         Success: "false",
         Data: struct{}{},
      })
   }

   user := &model.User{
      Password: &request.Password,
      Email: request.Email,
      Status: 1,
      CreatedBy: request.Email,
   }

   user = a.UserService.CreateUser(user)
   response := model.ResponseJSON{
      Message: "Register User Success",
      Success: "true",
      Data: user,
   }

   return c.JSON(response)
}