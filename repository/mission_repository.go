package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/shoelfikar/finpay-realtime-transaction/model"
)

type missionRepository struct {
	DB *sql.DB
}

type MissionRepository interface {
	CreateMission(mission model.Missions) model.Missions
	GetAllMission() []*model.Missions
	GetMissionByID(missionId int) *model.Missions
}

func NewMissionRepository(db *sql.DB) MissionRepository {
	return &missionRepository{
		DB: db,
	}
}

func (m *missionRepository) CreateMission(mission model.Missions) model.Missions {
	tx, err := m.DB.Begin()
	if err != nil {
		panic("Error database transaction"+ err.Error())
	}
	ctx := context.Background()

	query := `
		INSERT INTO missions (name, type, condition, point, status, created_by)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	_, err = tx.ExecContext(ctx, query, &mission.Name, &mission.Type, &mission.Condition, &mission.Point, &mission.Status, &mission.CreatedBy)

	if err != nil {
		tx.Rollback()
		panic("error insert new mission" + err.Error())
	}

	tx.Commit()
	return mission
}

func (m *missionRepository) GetMissionByID(missionId int) *model.Missions {
	ctx := context.Background()
	var mission model.Missions
	query := `
		SELECT id, name, type, condition, point, status, created_at, updated_at, created_by
		FROM missions WHERE id = $1
	`

	result := m.DB.QueryRowContext(ctx, query, &missionId).Scan(&mission.Id, &mission.Name,  &mission.Type, &mission.Condition, &mission.Point, &mission.Status, &mission.CreatedAt, &mission.UpdatedAt, &mission.CreatedBy)

	if result == sql.ErrNoRows {
		return nil
	}

	return &mission
}

func (m *missionRepository) GetAllMission() []*model.Missions {
	ctx := context.Background()
	var (
		
		missions []*model.Missions
		condition []byte
	)

	query := `
		SELECT id, name, type, condition, point, status, created_at, updated_at, created_by
		FROM missions order by id desc
	`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil
	}

	for rows.Next() {
		mission := model.Missions{}
		rows.Scan(&mission.Id, &mission.Name, &mission.Type, &condition, &mission.Point, &mission.Status, &mission.CreatedAt, &mission.UpdatedAt, &mission.CreatedBy)

		_ = json.Unmarshal(condition, &mission.Condition)


		missions = append(missions, &mission)
	}

	return missions
}