package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEventType(t *testing.T) {

	et, err := NewEventType("1", "Login")

	assert.Nil(t, err)
	assert.NotNil(t, et)
	assert.Equal(t, et.Code, "1")
	assert.Equal(t, et.Description, "Login")
	assert.NotEmpty(t, et.Id)
	assert.NotEmpty(t, et.CreatedAt)

}

func TestNewEventType_WhenCodeIsRequired(t *testing.T) {

	_, err := NewEventType("", "Login")

	assert.NotNil(t, err)
	assert.Equal(t, err, ErrCodeIsRequired)

}

func TestNewEventType_WhenDescriptionIsRequired(t *testing.T) {

	_, err := NewEventType("1", "")

	assert.NotNil(t, err)
	assert.Equal(t, err, ErrDescriptionIsRequired)

}
