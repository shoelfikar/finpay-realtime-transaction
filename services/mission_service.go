package services

import (
	"github.com/shoelfikar/finpay-realtime-transaction/model"
	"github.com/shoelfikar/finpay-realtime-transaction/repository"
)

type missionService struct {
	MissionRepository repository.MissionRepository
}

type MissionService interface {
	CreateMission(mission model.Mission) model.Mission
}

func NewMissionService(mission repository.MissionRepository) MissionService {
	return &missionService{
		MissionRepository: mission,
	}
}

func (m *missionService) CreateMission(mission model.Mission) model.Mission {
	return mission
}