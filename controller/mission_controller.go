package controller

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/shoelfikar/finpay-realtime-transaction/model"
	"github.com/shoelfikar/finpay-realtime-transaction/services"
	"github.com/shoelfikar/finpay-realtime-transaction/utils"
)

type missionController struct {
	MissionService services.MissionService
	validate       *validator.Validate
}

type MissionController interface {
	CreateMission(c *fiber.Ctx) error
	GetAllMission(c *fiber.Ctx) error
}

func NewMissionController(mission services.MissionService, validate *validator.Validate) MissionController {
	return &missionController{
		MissionService: mission,
		validate:       validate,
	}
}

// Missions godoc
// @Summary Create User Mission
// @Description Create User Mission
// @Tags Mission
// @Accept  json
// @Produce  json
// @Security  BearerAuth
// @Param request body model.MissionRequest true "Mission"
// @Router /mission/create [post]
func (m *missionController) CreateMission(c *fiber.Ctx) error {
	body := c.Body()
	var (
		request model.MissionRequest
	)
	err := json.Unmarshal(body, &request)
	if err != nil {
		return err
	}

	err = m.validate.Struct(request)
	if ok := utils.ValidationErrors(err); ok {
		message := utils.GetErrorMessagevalidator(err)
		return c.Status(fiber.StatusBadRequest).JSON(model.ResponseJSON{
			Success:         "false",
			Message:         "validation error",
			ValidationError: message,
		})
	}

	condition, _ := json.Marshal(request.Condition)

	mission := model.Missions{
		Name: request.Name,
		Type: request.Type,
		Condition: condition,
		Point: request.Point,
		Status: true,
		CreatedBy: "system",
	}

	result := m.MissionService.CreateMission(mission)

	response := model.ResponseJSON{
		Message: "Login Success",
		Success: "true",
		Data: result,
	}

	return c.JSON(response)
}

// Missions godoc
// @Summary Get All User Mission
// @Description Get All User Mission
// @Tags Mission
// @Accept  json
// @Produce  json
// @Security  BearerAuth
// @Router /mission/all [get]
func (m *missionController) GetAllMission(c *fiber.Ctx) error {
	missions := m.MissionService.GetAllMission()
	response := model.ResponseJSON{
		Message: "Get all mission",
		Success: "true",
		Data: missions,
	}

	return c.JSON(response)
}