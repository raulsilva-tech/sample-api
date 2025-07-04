package entity

import "errors"

var (
	ErrNameIsRequired        = errors.New("name is required")
	ErrEmailIsRequired       = errors.New("email is required")
	ErrPasswordIsRequired    = errors.New("password is required")
	ErrEmailIsInvalid        = errors.New("email is invalid")
	ErrCodeIsRequired        = errors.New("code is required")
	ErrDescriptionIsRequired = errors.New("description is required")
	ErrEventTypeIsRequired   = errors.New("event type is required")
	ErrUserIdIsRequired      = errors.New("user id is required")
)
