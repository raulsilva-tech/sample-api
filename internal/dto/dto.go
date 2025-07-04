package dto

type UserInput struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type EventTypeInput struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type EventInput struct {
	EventTypeId string `json:"event_type_id"`
	UserId      string `json:"user_id"`
}
