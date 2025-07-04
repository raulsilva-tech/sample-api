package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/raulsilva-tech/SampleAPI/internal/db"
	"github.com/raulsilva-tech/SampleAPI/internal/entity"
)

type EventTypeRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewEventTypeRepository(dbConn *sql.DB) *EventTypeRepository {
	return &EventTypeRepository{
		DB:      dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *EventTypeRepository) Insert(ctx context.Context, eventType entity.EventType) error {
	return r.Queries.CreateEventType(ctx, db.CreateEventTypeParams{
		ID:          eventType.Id.String(),
		Code:        eventType.Code,
		Description: sql.NullString{String: eventType.Description, Valid: true},
		CreatedAt:   eventType.CreatedAt,
		UpdatedAt:   sql.NullTime{Time: eventType.UpdatedAt, Valid: false},
	})
}

func (r *EventTypeRepository) Update(ctx context.Context, eventType entity.EventType) error {
	return r.Queries.UpdateEventType(ctx, db.UpdateEventTypeParams{
		ID:          eventType.Id.String(),
		Code:        eventType.Code,
		Description: sql.NullString{String: eventType.Description, Valid: true},
		UpdatedAt:   sql.NullTime{Time: eventType.UpdatedAt, Valid: true},
	})
}

func (r *EventTypeRepository) Delete(ctx context.Context, id string) error {
	return r.Queries.DeleteEventType(ctx, id)
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

func (r *EventTypeRepository) GetAll(ctx context.Context) ([]*entity.EventType, error) {

	dbuList, err := r.Queries.ListEventTypes(ctx)
	if err != nil {
		return nil, err
	}

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
	}

	return eventTypeList, nil

}
