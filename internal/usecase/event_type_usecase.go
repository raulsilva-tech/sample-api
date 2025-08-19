package usecase

import (
	"context"
	"log"
	"time"

	"github.com/raulsilva-tech/SampleAPI/internal/entity"
	"github.com/raulsilva-tech/SampleAPI/internal/repository"
)

type EventTypeUseCaseInput struct {
	Id          string
	Code        string
	Description string
	UserId      string
}

type EventTypeUseCaseOutput struct {
	Id          string
	Code        string
	Description string
	CreatedAt   time.Time
}

type EventTypeUseCase struct {
	EventTypeRepository *repository.EventTypeRepository
	EventRepository     *repository.EventRepository
}

func NewEventTypeUseCase(etr *repository.EventTypeRepository, er *repository.EventRepository) *EventTypeUseCase {
	return &EventTypeUseCase{
		EventTypeRepository: etr,
		EventRepository:     er,
	}
}

func (uc *EventTypeUseCase) RegisterEventType(ctx context.Context, input EventTypeUseCaseInput) (EventTypeUseCaseOutput, error) {

	evType, err := entity.NewEventType(input.Code, input.Description)
	if err != nil {
		return EventTypeUseCaseOutput{}, err
	}
	err = uc.EventTypeRepository.Insert(ctx, *evType)
	if err != nil {
		return EventTypeUseCaseOutput{}, err
	}

	//inserting event to register event type creation
	event, err := entity.NewEvent(uc.EventTypeRepository.EventTypeMap[entity.EventTypeInsert], input.UserId, entity.TableEventType, evType.Id.String())
	if err == nil {
		err = uc.EventRepository.Insert(ctx, *event)
		if err != nil {
			log.Println("Failed to insert 'insert' event", err)
		}
	} else {
		log.Println("Failed to create insert event", err)
	}

	return EventTypeUseCaseOutput{
		Id:          evType.Id.String(),
		Code:        evType.Code,
		Description: evType.Description,
		CreatedAt:   evType.CreatedAt,
	}, nil
}

func (uc *EventTypeUseCase) UpdateEventType(ctx context.Context, input EventTypeUseCaseInput) error {

	evType, err := entity.NewEventType(input.Code, input.Description)
	if err != nil {
		return err
	}
	err = uc.EventTypeRepository.Update(ctx, *evType)
	if err != nil {
		return err
	}

	//inserting event to register event type update
	event, err := entity.NewEvent(uc.EventTypeRepository.EventTypeMap[entity.EventTypeUpdate], input.UserId, entity.TableEventType, evType.Id.String())
	if err == nil {
		err = uc.EventRepository.Insert(ctx, *event)
		if err != nil {
			log.Println("Failed to insert 'update' event", err)
		}
	} else {
		log.Println("Failed to create 'update' event", err)
	}

	return err
}
func (uc *EventTypeUseCase) RemoveEventType(ctx context.Context, id string, userId string) error {

	err := uc.EventTypeRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	//inserting event to register event type update
	event, err := entity.NewEvent(uc.EventTypeRepository.EventTypeMap[entity.EventTypeDelete], userId, entity.TableEventType, id)
	if err == nil {
		err = uc.EventRepository.Insert(ctx, *event)
		if err != nil {
			log.Println("Failed to insert 'delete' event", err)
		}
	} else {
		log.Println("Failed to create 'delete' event", err)
	}

	return nil
}

func (uc *EventTypeUseCase) GetEventType(ctx context.Context, id string, userId string) (EventTypeUseCaseOutput, error) {

	evType, err := uc.EventTypeRepository.GetOne(ctx, id)
	if err != nil {
		return EventTypeUseCaseOutput{}, err
	}

	return EventTypeUseCaseOutput{
		Id:          evType.Id.String(),
		Code:        evType.Code,
		Description: evType.Description,
		CreatedAt:   evType.CreatedAt,
	}, nil
}

func (uc *EventTypeUseCase) GetAllEventTypes(ctx context.Context, userId string) ([]EventTypeUseCaseOutput, error) {

	evTypeList, err := uc.EventTypeRepository.GetAll(ctx)
	if err != nil {
		return []EventTypeUseCaseOutput{}, err
	}

	var outputList []EventTypeUseCaseOutput

	for _, evType := range evTypeList {

		output := EventTypeUseCaseOutput{
			Id:          evType.Id.String(),
			Code:        evType.Code,
			Description: evType.Description,
			CreatedAt:   evType.CreatedAt,
		}
		outputList = append(outputList, output)
	}

	return outputList, nil
}
