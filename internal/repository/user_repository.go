package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/raulsilva-tech/SampleAPI/internal/db"
	"github.com/raulsilva-tech/SampleAPI/internal/entity"
)

type UserRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewUserRepository(dbConn *sql.DB) *UserRepository {
	return &UserRepository{
		DB:      dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *UserRepository) Insert(ctx context.Context, user entity.User) error {
	return r.Queries.CreateUser(ctx, db.CreateUserParams{
		ID:        user.Id.String(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: sql.NullTime{Time: user.UpdatedAt, Valid: false},
	})
}

func (r *UserRepository) Update(ctx context.Context, user entity.User) error {
	return r.Queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:        user.Id.String(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		UpdatedAt: sql.NullTime{Time: user.UpdatedAt, Valid: true},
	})
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	return r.Queries.DeleteUser(ctx, id)
}

func (r *UserRepository) GetOne(ctx context.Context, id string) (*entity.User, error) {
	dbu, err := r.Queries.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	uuidId, err := uuid.Parse(dbu.ID)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		Id:        uuidId,
		Name:      dbu.Name,
		Email:     dbu.Email,
		Password:  dbu.Password,
		CreatedAt: dbu.CreatedAt,
		UpdatedAt: dbu.UpdatedAt.Time,
	}, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*entity.User, error) {

	dbuList, err := r.Queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var userList []*entity.User
	for _, dbu := range dbuList {

		uuidId, err := uuid.Parse(dbu.ID)
		if err != nil {
			return nil, err
		}
		userList = append(userList, &entity.User{
			Id:        uuidId,
			Name:      dbu.Name,
			Email:     dbu.Email,
			Password:  dbu.Password,
			CreatedAt: dbu.CreatedAt,
			UpdatedAt: dbu.UpdatedAt.Time,
		})
	}

	return userList, nil

}
