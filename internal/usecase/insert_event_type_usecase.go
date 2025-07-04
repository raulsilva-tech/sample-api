package usecase

import (
	"context"
	"time"

	"github.com/raulsilva-tech/SampleAPI/internal/entity"
	"github.com/raulsilva-tech/SampleAPI/internal/repository"
)

type InsertEventTypeUseCaseInput struct {
	Id          string
	Code        string
	Description string
	UserId      string
}

type InsertEventTypeUseCaseOutput struct {
	Id          string
	Code        string
	Description string
	CreatedAt   time.Time
}

type InsertEventTypeUseCase struct {
	EventTypeRepository *repository.EventTypeRepository
	EventRepository     *repository.EventRepository
}

func NewInsertEventTypeUseCase(etr *repository.EventTypeRepository, er *repository.EventRepository) *InsertEventTypeUseCase {
	return &InsertEventTypeUseCase{
		EventTypeRepository: etr,
		EventRepository:     er,
	}
}

func (uc *InsertEventTypeUseCase) Execute(ctx context.Context, input InsertEventTypeUseCaseInput) error {

	evType, err := entity.NewEventType(input.Code, input.Description)
	if err != nil {
		return err
	}
	err = uc.EventTypeRepository.Insert(ctx, *evType)
	if err != nil {
		return err
	}

	//inserting event to register event type creation
	event, err := entity.NewEvent(evType.Id.String(), input.UserId)
	if err != nil {
		return err
	}
	err = uc.EventRepository.Insert(ctx, *event)

	return err
}
