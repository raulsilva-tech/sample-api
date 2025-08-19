package dto

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EventTypeInput struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type EventInput struct {
	EventTypeId string `json:"event_type_id"`
	UserId      string `json:"user_id"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
