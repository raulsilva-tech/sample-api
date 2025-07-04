package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/raulsilva-tech/SampleAPI/internal/db"
	"github.com/raulsilva-tech/SampleAPI/internal/entity"
)

type EventRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewEventRepository(dbConn *sql.DB) *EventRepository {
	return &EventRepository{
		DB:      dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *EventRepository) Insert(ctx context.Context, ev entity.Event) error {
	return r.Queries.CreateEvent(ctx, db.CreateEventParams{
		ID:          ev.Id.String(),
		EventTypeID: ev.EvType.Id.String(),
		UserID:      ev.EvUser.Id.String(),
		CreatedAt:   ev.CreatedAt,
	})
}

func (r *EventRepository) Update(ctx context.Context, ev entity.Event) error {
	return r.Queries.UpdateEvent(ctx, db.UpdateEventParams{
		ID:          ev.Id.String(),
		EventTypeID: ev.EvType.Id.String(),
		UserID:      ev.EvUser.Id.String(),
	})
}

func (r *EventRepository) Delete(ctx context.Context, id string) error {
	return r.Queries.DeleteEvent(ctx, id)
}

func (r *EventRepository) GetOne(ctx context.Context, id string) (*entity.Event, error) {
	ev, err := r.Queries.GetEvent(ctx, id)
	if err != nil {
		return nil, err
	}

	uuidId, err := uuid.Parse(ev.ID)
	if err != nil {
		return nil, err
	}
	uuidUserId, err := uuid.Parse(ev.UserID)
	if err != nil {
		return nil, err
	}
	uuidEventTypeId, err := uuid.Parse(ev.EventTypeID)
	if err != nil {
		return nil, err
	}
	return &entity.Event{
		Id:        uuidId,
		EvType:    entity.EventType{Id: uuidEventTypeId},
		EvUser:    entity.User{Id: uuidUserId},
		CreatedAt: ev.CreatedAt,
	}, nil
}

func (r *EventRepository) GetAll(ctx context.Context) ([]*entity.Event, error) {

	evList, err := r.Queries.ListEvents(ctx)
	if err != nil {
		return nil, err
	}

	var eventList []*entity.Event
	for _, ev := range evList {

		uuidId, err := uuid.Parse(ev.ID)
		if err != nil {
			return nil, err
		}
		uuidUserId, err := uuid.Parse(ev.UserID)
		if err != nil {
			return nil, err
		}
		uuidEventTypeId, err := uuid.Parse(ev.EventTypeID)
		if err != nil {
			return nil, err
		}

		eventList = append(eventList, &entity.Event{
			Id:        uuidId,
			EvType:    entity.EventType{Id: uuidEventTypeId},
			EvUser:    entity.User{Id: uuidUserId},
			CreatedAt: ev.CreatedAt,
		})
	}

	return eventList, nil

}
