package services

import (
	"github.com/shoelfikar/finpay-realtime-transaction/model"
	"github.com/shoelfikar/finpay-realtime-transaction/repository"
)

type missionService struct {
	MissionRepository repository.MissionRepository
}

type MissionService interface {
	CreateMission(mission model.Missions) model.Missions
	GetMissionByID(missionId int) model.Missions
	GetAllMission() []*model.Missions
}

func NewMissionService(mission repository.MissionRepository) MissionService {
	return &missionService{
		MissionRepository: mission,
	}
}

func (m *missionService) CreateMission(mission model.Missions) model.Missions {
	result := m.MissionRepository.CreateMission(mission)
	return result
}

func (m *missionService) GetMissionByID(missionId int) model.Missions {
	mission := m.MissionRepository.GetMissionByID(missionId)
	return *mission
}

func (m *missionService) GetAllMission() []*model.Missions {
	missions := m.MissionRepository.GetAllMission()
	return missions
}