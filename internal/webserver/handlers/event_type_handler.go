package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulsilva-tech/SampleAPI/internal/dto"
	"github.com/raulsilva-tech/SampleAPI/internal/repository"
	"github.com/raulsilva-tech/SampleAPI/internal/usecase"
)

type EventTypeHandler struct {
	EventTypeRepository *repository.EventTypeRepository
	EventRepository     *repository.EventRepository
}

func NewEventTypeHandler(etr *repository.EventTypeRepository, er *repository.EventRepository) *EventTypeHandler {
	return &EventTypeHandler{
		EventTypeRepository: etr,
		EventRepository:     er,
	}
}

func (h *EventTypeHandler) Insert(c *gin.Context) {

	var input dto.EventTypeInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authenticated user. An user is required to insert records."})
		return
	}
	ucInput := usecase.InsertEventTypeUseCaseInput{
		Code:        input.Code,
		Description: input.Description,
		UserId:      userID.(string),
	}

	uc := usecase.NewInsertEventTypeUseCase(h.EventTypeRepository, h.EventRepository)
	err = uc.Execute(c.Request.Context(), ucInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, map[string]any{"msg": "created successfully"})
}

func (h *EventTypeHandler) Delete(r *gin.Context) {

	id := r.Param("id")
	if id == "" {
		r.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	// uc := usecase.NewDeleteEventTypeUseCase(h.EventTypeRepository, h.EventRepository)
	// err := uc.Execute(id)
	// if err != nil {
	// 	r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// return
	// }

	r.JSON(http.StatusOK, map[string]any{"msg": "deleted successfully"})
}
