package usecase

import (
	"context"
	"time"

	"github.com/raulsilva-tech/SampleAPI/internal/entity"
	"github.com/raulsilva-tech/SampleAPI/internal/repository"
)

type EventUseCase struct {
	EventRepository *repository.EventRepository
}

func NewEventUseCase(er *repository.EventRepository) *EventUseCase {
	return &EventUseCase{
		EventRepository: er,
	}
}

type EventUseCaseInput struct {
	EventTypeId string
	UserId      string
	TargetTable string
	TargetId    string
}

type EventUseCaseOutput struct {
	Id          string
	EventTypeId string
	UserId      string
	TargetTable string
	TargetId    string
}

func (uc *EventUseCase) RegisterEvent(ctx context.Context, input EventUseCaseInput) (EventUseCaseOutput, error) {

	ev, err := entity.NewEvent(input.EventTypeId, input.UserId, input.TargetTable, input.TargetId, time.Now())
	if err != nil {
		return EventUseCaseOutput{}, err
	}
	err = uc.EventRepository.Insert(ctx, *ev)
	if err != nil {
		return EventUseCaseOutput{}, err
	}

	return EventUseCaseOutput{
		Id:          ev.Id.String(),
		EventTypeId: ev.EvType.Id.String(),
		UserId:      ev.EvUser.Id.String(),
		TargetTable: ev.TargetTable,
		TargetId:    ev.TargetId,
	}, nil
}
