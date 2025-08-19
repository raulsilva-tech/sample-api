package usecase

import (
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raulsilva-tech/SampleAPI/internal/entity"
	"github.com/raulsilva-tech/SampleAPI/internal/repository"
)

type LoginUseCaseInput struct {
	Email        string
	Password     string
	JWTSecret    string
	JWTExpiresIn int
}

type LoginUseCaseOutput struct {
	Id        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Token     string
}

type UserUseCase struct {
	UserRepository      *repository.UserRepository
	EventRepository     *repository.EventRepository
	EventTypeRepository *repository.EventTypeRepository
}

type UserUseCaseInput struct {
	Name          string
	Email         string
	Password      string
	CreatorUserId string
}

type UserUseCaseOutput struct {
	Id        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func NewUserUseCase(ur *repository.UserRepository, er *repository.EventRepository, etr *repository.EventTypeRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository:      ur,
		EventRepository:     er,
		EventTypeRepository: etr,
	}

}

func (uc *UserUseCase) RegisterUser(ctx context.Context, input UserUseCaseInput) (UserUseCaseOutput, error) {

	u, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return UserUseCaseOutput{}, err
	}

	err = uc.UserRepository.Insert(ctx, *u)
	if err != nil {
		return UserUseCaseOutput{}, err
	}

	ev, err := entity.NewEvent(uc.EventTypeRepository.EventTypeMap[entity.EventTypeInsert], input.CreatorUserId, entity.TableUser, u.Id.String())
	if err == nil {

		err = uc.EventRepository.Insert(ctx, *ev)
		if err != nil {
			log.Println("failed to insert 'insert' event")
		}
	} else {
		log.Println("failed to create 'insert' event")
	}

	return UserUseCaseOutput{
		Id:        u.Id.String(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}, nil

}

func (uc *UserUseCase) Login(ctx context.Context, input LoginUseCaseInput) (LoginUseCaseOutput, error) {

	user, err := uc.UserRepository.Login(ctx, entity.User{Email: input.Email, Password: input.Password})
	if err != nil {

		return LoginUseCaseOutput{}, err
	}

	//inserting event to register login
	event, err := entity.NewEvent(uc.EventTypeRepository.EventTypeMap[entity.EventTypeLogin], user.Id.String(), "", "")
	if err == nil {

		err = uc.EventRepository.Insert(ctx, *event)
		if err != nil {
			log.Println("failed to insert login event ", err)
		}
	} else {
		log.Println("failed to create login event ", err)
	}

	token, err := generateToken(user.Id.String(), input.JWTSecret, input.JWTExpiresIn)
	if err != nil {
		return LoginUseCaseOutput{}, err
	}

	return LoginUseCaseOutput{
		Id:        user.Id.String(),
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     token,
	}, nil

}

func generateToken(userId string, jwtSecret string, jwtExpiresIn int) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Minute * time.Duration(jwtExpiresIn)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// type Claims struct{
// 	UserId string `json:"user_id"`
// 	jwt.RegisteredClaims
// }
