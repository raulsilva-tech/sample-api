package usecase

import "github.com/raulsilva-tech/SampleAPI/internal/repository"

type InsertUserEventUsecase struct {
	UserRepository  repository.UserRepository
	EventRepository repository.EventRepository
}

func NewInsertUserEventUseCase(ur repository.UserRepository, er repository.EventRepository) *InsertUserEventUsecase {
	return &InsertUserEventUsecase{
		UserRepository:  ur,
		EventRepository: er,
	}
}

func (uc *InsertUserEventUsecase) Execute() error {

	return nil
}
