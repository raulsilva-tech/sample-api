package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	EvType      EventType `json:"type"`
	EvUser      User      `json:"user"`
	TargetTable string
	TargetId    string
}

func NewEvent(eventTypeId, userId, targetTable, targetId string) (*Event, error) {

	uuidEvTypeId, err := uuid.Parse(eventTypeId)
	if err != nil {
		return nil, errors.New("event_type_id: invalid uuid string > " + eventTypeId)
	}
	uuidUserId, err := uuid.Parse(userId)
	if err != nil {
		return nil, errors.New("user_id: invalid uuid string > " + userId)
	}

	e := &Event{
		Id:          uuid.New(),
		CreatedAt:   time.Now(),
		EvType:      EventType{Id: uuidEvTypeId},
		EvUser:      User{Id: uuidUserId},
		TargetTable: targetTable,
		TargetId:    targetId,
	}
	return e, nil
}
