package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/raulsilva-tech/SampleAPI/internal/entity"
	"github.com/raulsilva-tech/SampleAPI/internal/repository"
	"golang.org/x/crypto/bcrypt"
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
	Id            string
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
	UpdatedAt time.Time
}

func NewUserUseCase(ur *repository.UserRepository, er *repository.EventRepository, etr *repository.EventTypeRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository:      ur,
		EventRepository:     er,
		EventTypeRepository: etr,
	}

}

func (uc *UserUseCase) RemoveUser(ctx context.Context, id string, deleterUserId string) error {

	err := uc.UserRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	ev, err := entity.NewEvent(uc.EventTypeRepository.EventTypeMap[entity.EventTypeDelete], deleterUserId, entity.TableUser, id, time.Now())
	if err == nil {
		err = uc.EventRepository.Insert(ctx, *ev)
		if err != nil {
			log.Println("failed to insert 'delete' event:", err.Error())
		}
	} else {
		log.Println("failed to create 'delete' event:", err.Error())
	}

	return nil
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

	ev, err := entity.NewEvent(uc.EventTypeRepository.EventTypeMap[entity.EventTypeInsert], input.CreatorUserId, entity.TableUser, u.Id.String(), time.Now())
	if err == nil {
		err = uc.EventRepository.Insert(ctx, *ev)
		if err != nil {
			log.Println("failed to insert 'insert' event: ", err.Error())
		}
	} else {
		log.Println("failed to create 'insert' event: ", err.Error())
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
	event, err := entity.NewEvent(uc.EventTypeRepository.EventTypeMap[entity.EventTypeLogin], user.Id.String(), "", "", time.Now())
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

func (uc *UserUseCase) UpdateUser(ctx context.Context, input UserUseCaseInput) error {

	uuidUserId, err := uuid.Parse(input.Id)
	if err != nil {
		return errors.New("user_id: invalid uuid string > " + input.Id)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	input.Password = string(hash)

	u := entity.User{
		Id:        uuidUserId,
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		UpdatedAt: time.Now(),
	}
	err = uc.UserRepository.Update(ctx, u)
	if err != nil {
		return err
	}

	ev, err := entity.NewEvent(uc.EventTypeRepository.EventTypeMap[entity.EventTypeUpdate], input.CreatorUserId, entity.TableUser, u.Id.String(), time.Now())
	if err == nil {
		err = uc.EventRepository.Insert(ctx, *ev)
		if err != nil {
			log.Println("failed to insert 'insert' event: ", err.Error())
		}
	} else {
		log.Println("failed to create 'insert' event: ", err.Error())
	}

	return nil

}

func (uc *UserUseCase) GetUser(ctx context.Context, id string) (UserUseCaseOutput, error) {

	user, err := uc.UserRepository.GetOne(ctx, id)
	if err != nil {
		return UserUseCaseOutput{}, err
	}

	return UserUseCaseOutput{
		Id:        user.Id.String(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (uc *UserUseCase) GetAll(ctx context.Context) ([]UserUseCaseOutput, error) {

	userList, err := uc.UserRepository.GetAll(ctx)
	if err != nil {
		return []UserUseCaseOutput{}, err
	}

	var outputList []UserUseCaseOutput

	for _, u := range userList {

		var item UserUseCaseOutput
		item.Id = u.Id.String()
		item.Name = u.Name
		item.CreatedAt = u.CreatedAt
		item.Email = u.Email
		item.Password = u.Password
		item.UpdatedAt = u.UpdatedAt
		outputList = append(outputList, item)
	}

	return outputList, nil
}
