package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	EvType    EventType `json:"type"`
	EvUser    User      `json:"user"`
}

func NewEvent(eventTypeId, userId string) (*Event, error) {

	uuidEvTypeId, err := uuid.Parse(eventTypeId)
	if err != nil {
		return nil, errors.New("event_type_id: invalid uuid string")
	}
	uuidUserId, err := uuid.Parse(userId)
	if err != nil {
		return nil, errors.New("user_id: invalid uuid string")
	}

	e := &Event{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		EvType:    EventType{Id: uuidEvTypeId},
		EvUser:    User{Id: uuidUserId},
	}
	return e, nil
}
