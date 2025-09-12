package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/raulsilva-tech/sample-api/internal/db"
	"github.com/raulsilva-tech/sample-api/internal/entity"
)

type EventTypeRepository struct {
	DB           *sql.DB
	Queries      *db.Queries
	EventTypeMap map[string]string
}

func NewEventTypeRepository(ctx context.Context, dbConn *sql.DB) (*EventTypeRepository, error) {

	repo := &EventTypeRepository{
		DB:      dbConn,
		Queries: db.New(dbConn),
	}

	_, err := repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *EventTypeRepository) Insert(ctx context.Context, eventType entity.EventType) error {
	err := r.Queries.CreateEventType(ctx, db.CreateEventTypeParams{
		ID:          eventType.Id.String(),
		Code:        eventType.Code,
		Description: sql.NullString{String: eventType.Description, Valid: true},
		CreatedAt:   eventType.CreatedAt,
		UpdatedAt:   sql.NullTime{Time: eventType.UpdatedAt, Valid: false},
	})

	if err != nil {
		return err
	}

	r.EventTypeMap[eventType.Code] = eventType.Id.String()

	return nil
}

func (r *EventTypeRepository) Update(ctx context.Context, eventType entity.EventType) error {
	err := r.Queries.UpdateEventType(ctx, db.UpdateEventTypeParams{
		ID:          eventType.Id.String(),
		Code:        eventType.Code,
		Description: sql.NullString{String: eventType.Description, Valid: true},
		UpdatedAt:   sql.NullTime{Time: eventType.UpdatedAt, Valid: true},
	})
	if err != nil {
		return err
	}
	r.EventTypeMap[eventType.Code] = eventType.Id.String()

	return nil
}

func (r *EventTypeRepository) Delete(ctx context.Context, id string) error {
	err := r.Queries.DeleteEventType(ctx, id)
	if err != nil {
		return err
	}
	for k, v := range r.EventTypeMap {
		if v == id {
			delete(r.EventTypeMap, k)
		}
	}

	return nil

}

func (r *EventTypeRepository) GetOne(ctx context.Context, id string) (*entity.EventType, error) {
	dbet, err := r.Queries.GetEventType(ctx, id)
	if err != nil {
		return nil, err
	}

	uuidId, err := uuid.Parse(dbet.ID)
	if err != nil {
		return nil, err
	}

	return &entity.EventType{
		Id:          uuidId,
		Code:        dbet.Code,
		Description: dbet.Description.String,
		CreatedAt:   dbet.CreatedAt,
		UpdatedAt:   dbet.UpdatedAt.Time,
	}, nil
}

func (r *EventTypeRepository) GetOneByCode(ctx context.Context, code string) (*entity.EventType, error) {
	dbet, err := r.Queries.GetEventTypeByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	uuidId, err := uuid.Parse(dbet.ID)
	if err != nil {
		return nil, err
	}

	return &entity.EventType{
		Id:          uuidId,
		Code:        dbet.Code,
		Description: dbet.Description.String,
		CreatedAt:   dbet.CreatedAt,
		UpdatedAt:   dbet.UpdatedAt.Time,
	}, nil
}

func (r *EventTypeRepository) GetAll(ctx context.Context) ([]*entity.EventType, error) {

	dbuList, err := r.Queries.ListEventTypes(ctx)
	if err != nil {
		return nil, err
	}

	r.EventTypeMap = make(map[string]string)

	var eventTypeList []*entity.EventType
	for _, dbet := range dbuList {

		uuidId, err := uuid.Parse(dbet.ID)
		if err != nil {
			return nil, err
		}
		eventTypeList = append(eventTypeList, &entity.EventType{
			Id:          uuidId,
			Code:        dbet.Code,
			Description: dbet.Description.String,
			CreatedAt:   dbet.CreatedAt,
			UpdatedAt:   dbet.UpdatedAt.Time,
		})
		r.EventTypeMap[dbet.Code] = dbet.ID
	}

	return eventTypeList, nil

}
