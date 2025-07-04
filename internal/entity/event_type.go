package entity

import (
	"time"

	"github.com/google/uuid"
)

type EventType struct {
	Id          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewEventType(code, description string) (*EventType, error) {

	et := &EventType{
		Id:          uuid.New(),
		Code:        code,
		Description: description,
		CreatedAt:   time.Now(),
	}

	return et, et.CheckFields()

}

func (et *EventType) CheckFields() error {

	if et.Code == "" {
		return ErrCodeIsRequired
	}
	if et.Description == "" {
		return ErrDescriptionIsRequired
	}

	return nil
}
