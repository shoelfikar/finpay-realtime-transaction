package repository

import (
	"database/sql"

	"github.com/shoelfikar/finpay-realtime-transaction/model"
)

type missionRepository struct {
	DB *sql.DB
}

type MissionRepository interface {
	CreateMission(mission model.Mission) model.Mission
}

func NewMissionRepository(db *sql.DB) MissionRepository {
	return &missionRepository{
		DB: db,
	}
}

func (m *missionRepository) CreateMission(mission model.Mission) model.Mission {
	
	return mission
}