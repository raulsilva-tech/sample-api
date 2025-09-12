package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/raulsilva-tech/sample-api/internal/entity"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	DB *sql.DB
	suite.Suite
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	suite.DB.Close()
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	dbConn, err := migrateDB()
	suite.NoError(err)
	suite.DB = dbConn
}

func (suite *UserRepositoryTestSuite) TestInsert() {

	ctx := context.Background()

	u, err := entity.NewUser("Raul", "raul_insert@gmail.com", "teste")
	suite.NoError(err)

	r := NewUserRepository(suite.DB)
	err = r.Insert(ctx, *u)
	suite.NoError(err)

}

func (suite *UserRepositoryTestSuite) TestUpdate() {
	ctx := context.Background()

	u, err := entity.NewUser("Raul", "raul_update@gmail.com", "teste")
	suite.NoError(err)

	r := NewUserRepository(suite.DB)
	err = r.Insert(ctx, *u)
	suite.NoError(err)

	time.Sleep(2 * time.Second)
	u.Name = "Renato"
	u.UpdatedAt = time.Now()
	err = r.Update(ctx, *u)
	suite.NoError(err)

	user, err := r.GetOne(ctx, u.Id.String())
	suite.NoError(err)
	suite.Equal(user.Name, u.Name)

}

func (suite *UserRepositoryTestSuite) TestGetOne() {
	ctx := context.Background()

	u, err := entity.NewUser("Raul", "raul_getOne@gmail.com", "teste")
	suite.NoError(err)

	r := NewUserRepository(suite.DB)
	err = r.Insert(ctx, *u)
	suite.NoError(err)
	user, err := r.GetOne(ctx, u.Id.String())
	suite.NoError(err)
	suite.Equal(user.Name, u.Name)

}

func (suite *UserRepositoryTestSuite) TestLogin() {
	ctx := context.Background()

	u, err := entity.NewUser("Raul", "raulLogin@gmail.com", "test")
	suite.NoError(err)
	fmt.Println(u.Id.String())
	r := NewUserRepository(suite.DB)
	err = r.Insert(ctx, *u)
	suite.NoError(err)
	user, err := r.Login(ctx, *u)
	suite.NoError(err)
	suite.Equal(user.Name, u.Name)

}
func (suite *UserRepositoryTestSuite) TestDelete() {
	ctx := context.Background()

	u, err := entity.NewUser("Raul", "raul_getOne@gmail.com", "teste")
	suite.NoError(err)

	r := NewUserRepository(suite.DB)
	err = r.Insert(ctx, *u)
	suite.NoError(err)
	err = r.Delete(ctx, u.Id.String())
	suite.NoError(err)

}

func (suite *UserRepositoryTestSuite) TestGetAll() {
	ctx := context.Background()

	u, err := entity.NewUser("Raul", "raul_getALL@gmail.com", "teste")
	suite.NoError(err)

	r := NewUserRepository(suite.DB)
	err = r.Insert(ctx, *u)
	suite.NoError(err)

	userList, err := r.GetAll(ctx)
	suite.NoError(err)
	suite.NotEmpty(userList)
	suite.Equal(len(userList), 1)

}

func migrateDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME
);
CREATE TABLE event_types (
    id CHAR(36) PRIMARY KEY,
    code VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME
);
CREATE TABLE events (
    id CHAR(36) PRIMARY KEY,
    created_at DATETIME NOT NULL,
    user_id char(36) NOT NULL,
    event_type_id char(36) NOT NULL,
    FOREIGN KEY (event_type_id) REFERENCES event_types(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)`)

	return db, err
}
