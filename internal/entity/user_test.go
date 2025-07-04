package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	//Arrange, act
	u, err := NewUser("Raul", "raul@gmail.com", "test")

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u.Name, "Raul")
	assert.Equal(t, u.Email, "raul@gmail.com")
	assert.NotEmpty(t, u.Password)
	assert.Equal(t, u.isPasswordValid("test"), true)
}

func TestNewUser_WhenNameIsRequired(t *testing.T) {
	_, err := NewUser("", "raul@gmail.com", "test")
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNameIsRequired)
}

func TestNewUser_WhenEmailIsRequired(t *testing.T) {
	_, err := NewUser("Raul", "", "test")
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrEmailIsRequired)
}
func TestNewUser_WhenPasswordIsRequired(t *testing.T) {
	_, err := NewUser("Raul", "ra@gmail.com", "")

	assert.NotNil(t, err)
	assert.Equal(t, err, ErrPasswordIsRequired)
}

func TestNewUser_WhenEmailIsInvalid(t *testing.T) {
	_, err := NewUser("Raul", "raul-Mail.com", "test")
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrEmailIsInvalid)
}
