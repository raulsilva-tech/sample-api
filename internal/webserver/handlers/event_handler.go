package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulsilva-tech/SampleAPI/internal/dto"
	"github.com/raulsilva-tech/SampleAPI/internal/repository"
	"github.com/raulsilva-tech/SampleAPI/internal/usecase"
)

type EventHandler struct {
	EventRepository *repository.EventRepository
}

func NewEventHandler(er *repository.EventRepository) *EventHandler {
	return &EventHandler{
		EventRepository: er,
	}
}

func (h *EventHandler) Insert(c *gin.Context) {

	_, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authenticated user"})
		return
	}

	var input dto.EventInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ucInput := usecase.EventUseCaseInput{
		EventTypeId: input.EventTypeId,
		UserId:      input.UserId,
		TargetTable: input.TargetTable,
		TargetId:    input.TargetId,
	}

	uc := usecase.NewEventUseCase(h.EventRepository)
	output, err := uc.RegisterEvent(c.Request.Context(), ucInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}
