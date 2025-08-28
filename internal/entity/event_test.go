package entity

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {

	userId := uuid.NewString()
	evTypeId := uuid.NewString()
	ev, err := NewEvent(evTypeId, userId, "", "", time.Now())

	assert.Nil(t, err)
	assert.NotNil(t, ev)
	assert.NotEmpty(t, ev.CreatedAt)
	assert.NotEmpty(t, ev.Id)
	assert.NotEmpty(t, ev.EvType.Id)
	assert.NotEmpty(t, ev.EvUser.Id)
	assert.Equal(t, ev.EvType.Id.String(), evTypeId)
	assert.Equal(t, ev.EvUser.Id.String(), userId)
}

func TestNewEvent_WhenEventTypeIsInvalid(t *testing.T) {

	_, err := NewEvent("dsd", uuid.NewString(), "", "", time.Now())

	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("event_type_id: invalid uuid string"))
}

func TestNewEvent_WhenUserIsInvalid(t *testing.T) {

	_, err := NewEvent(uuid.NewString(), "", "", "", time.Now())

	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("user_id: invalid uuid string"))
}
