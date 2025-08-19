package usecase

import "github.com/raulsilva-tech/SampleAPI/internal/repository"

type EventUserCase struct {
	UserRepository  *repository.UserRepository
	EventRepository *repository.EventRepository
}

func NewEventUseCase(ur *repository.UserRepository, er *repository.EventRepository) *EventUserCase {
	return &EventUserCase{
		UserRepository:  ur,
		EventRepository: er,
	}
}

func (uc *EventUserCase) RegisterEvent() error {

	return nil
}
