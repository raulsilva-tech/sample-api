package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/raulsilva-tech/SampleAPI/internal/entity"
	"github.com/stretchr/testify/suite"
)

type EventTypeRepositoryTestSuite struct {
	suite.Suite
	DB *sql.DB
}

func TestEventTypeRepositorySuite(t *testing.T) {
	suite.Run(t, new(EventTypeRepositoryTestSuite))
}

func (suite *EventTypeRepositoryTestSuite) TearDownSuite() {
	suite.DB.Close()
}

func (suite *EventTypeRepositoryTestSuite) SetupSuite() {
	dbConn, err := migrateDB()
	suite.NoError(err)
	suite.DB = dbConn
}

func (suite *EventTypeRepositoryTestSuite) TestInsert() {
	ctx := context.Background()

	et, err := entity.NewEventType("1", "Login")
	suite.NoError(err)

	r := NewEventTypeRepository(suite.DB)
	err = r.Insert(ctx, *et)
	suite.NoError(err)

}

func (suite *EventTypeRepositoryTestSuite) TestUpdate() {
	ctx := context.Background()

	et, err := entity.NewEventType("1U", "Login")
	suite.NoError(err)

	r := NewEventTypeRepository(suite.DB)
	err = r.Insert(ctx, *et)
	suite.NoError(err)

	time.Sleep(2 * time.Second)
	et.Description = "Logout"
	et.UpdatedAt = time.Now()
	err = r.Update(ctx, *et)
	suite.NoError(err)

	eventType, err := r.GetOne(ctx, et.Id.String())
	suite.NoError(err)
	suite.Equal(et.Description, eventType.Description)

}

func (suite *EventTypeRepositoryTestSuite) TestGetOne() {
	ctx := context.Background()

	et, err := entity.NewEventType("1G", "Login")
	suite.NoError(err)

	r := NewEventTypeRepository(suite.DB)
	err = r.Insert(ctx, *et)
	suite.NoError(err)
	eventType, err := r.GetOne(ctx, et.Id.String())

	suite.NoError(err)
	suite.Equal(eventType.Description, et.Description)

}

func (suite *EventTypeRepositoryTestSuite) TestDelete() {
	ctx := context.Background()

	et, err := entity.NewEventType("1D", "Login")
	suite.NoError(err)

	r := NewEventTypeRepository(suite.DB)
	err = r.Insert(ctx, *et)
	suite.NoError(err)
	err = r.Delete(ctx, et.Id.String())
	suite.NoError(err)

}

func (suite *EventTypeRepositoryTestSuite) TestGetAll() {
	ctx := context.Background()

	et, err := entity.NewEventType("1ALL", "Login")
	suite.NoError(err)

	r := NewEventTypeRepository(suite.DB)
	err = r.Insert(ctx, *et)
	suite.NoError(err)

	eventTypeList, err := r.GetAll(context.Background())
	suite.NoError(err)
	suite.NotEmpty(eventTypeList)
	suite.Equal(len(eventTypeList), 1)

}
