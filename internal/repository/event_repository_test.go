package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/raulsilva-tech/SampleAPI/internal/entity"
	"github.com/stretchr/testify/suite"
)

type EventRepositoryTestSuite struct {
	suite.Suite
	DB *sql.DB
}

func TestEventRepositorySuite(t *testing.T) {
	suite.Run(t, new(EventRepositoryTestSuite))
}

func (suite *EventRepositoryTestSuite) TearDownSuite() {
	suite.DB.Close()
}

func (suite *EventRepositoryTestSuite) SetupSuite() {
	dbConn, err := migrateDB()
	suite.NoError(err)
	suite.DB = dbConn
}

func (suite *EventRepositoryTestSuite) TestInsert() {
	ctx := context.Background()

	userIdStr := uuid.NewString()
	etIdStr := uuid.NewString()
	ev, err := entity.NewEvent( etIdStr, userIdStr)
	suite.NoError(err)

	r := NewEventRepository(suite.DB)
	err = r.Insert(ctx, *ev)
	suite.NoError(err)

}

func (suite *EventRepositoryTestSuite) TestUpdate() {
	ctx := context.Background()

	userIdStr := uuid.NewString()
	etIdStr := uuid.NewString()
	ev, err := entity.NewEvent( etIdStr, userIdStr)
	suite.NoError(err)

	r := NewEventRepository(suite.DB)
	err = r.Insert(ctx, *ev)
	suite.NoError(err)

	time.Sleep(2 * time.Second)
	newId := uuid.New()
	ev.EvType.Id = newId
	err = r.Update(ctx, *ev)
	suite.NoError(err)

	event, err := r.GetOne(ctx, ev.Id.String())
	suite.NoError(err)
	suite.Equal(event.EvType.Id.String(), ev.EvType.Id.String())

}

func (suite *EventRepositoryTestSuite) TestGetOne() {
	ctx := context.Background()

	userIdStr := uuid.NewString()
	etIdStr := uuid.NewString()
	ev, err := entity.NewEvent( etIdStr, userIdStr)
	suite.NoError(err)

	r := NewEventRepository(suite.DB)
	err = r.Insert(ctx, *ev)
	suite.NoError(err)
	event, err := r.GetOne(ctx, ev.Id.String())

	suite.NoError(err)
	suite.Equal(event.EvUser.Id.String(), userIdStr)

}

func (suite *EventRepositoryTestSuite) TestDelete() {
	ctx := context.Background()

	userIdStr := uuid.NewString()
	etIdStr := uuid.NewString()
ev, err := entity.NewEvent( etIdStr, userIdStr)
	suite.NoError(err)

	r := NewEventRepository(suite.DB)
	err = r.Insert(ctx, *ev)
	suite.NoError(err)
	err = r.Delete(ctx, ev.Id.String())
	suite.NoError(err)

}

func (suite *EventRepositoryTestSuite) TestGetAll() {

	ctx := context.Background()

	userIdStr := uuid.NewString()
	etIdStr := uuid.NewString()
	ev, err := entity.NewEvent( etIdStr, userIdStr)
	suite.NoError(err)

	r := NewEventRepository(suite.DB)
	err = r.Insert(ctx, *ev)
	suite.NoError(err)

	eventList, err := r.GetAll(ctx)
	suite.NoError(err)
	suite.NotEmpty(eventList)
	suite.Equal(len(eventList), 1)

}
