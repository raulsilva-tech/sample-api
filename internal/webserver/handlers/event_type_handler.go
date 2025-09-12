package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulsilva-tech/sample-api/internal/dto"
	"github.com/raulsilva-tech/sample-api/internal/repository"
	"github.com/raulsilva-tech/sample-api/internal/usecase"
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

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authenticated user. An user is required to insert records."})
		return
	}

	var input dto.EventTypeInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ucInput := usecase.EventTypeUseCaseInput{
		Code:        input.Code,
		Description: input.Description,
		UserId:      userID.(string),
	}

	uc := usecase.NewEventTypeUseCase(h.EventTypeRepository, h.EventRepository)
	_, err = uc.RegisterEventType(c.Request.Context(), ucInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, map[string]any{"msg": "created successfully"})
}

func (h *EventTypeHandler) Update(c *gin.Context) {

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authenticated user. An user is required to update records."})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var input dto.EventTypeInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ucInput := usecase.EventTypeUseCaseInput{
		Id:          id,
		Code:        input.Code,
		Description: input.Description,
		UserId:      userID.(string),
	}

	uc := usecase.NewEventTypeUseCase(h.EventTypeRepository, h.EventRepository)
	err = uc.UpdateEventType(c.Request.Context(), ucInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]any{"msg": "updated successfully"})
}

func (h *EventTypeHandler) Delete(c *gin.Context) {

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authenticated user. An user is required to delete records."})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	uc := usecase.NewEventTypeUseCase(h.EventTypeRepository, h.EventRepository)
	err := uc.RemoveEventType(c.Request.Context(), id, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]any{"msg": "deleted successfully"})
}

func (h *EventTypeHandler) GetOne(c *gin.Context) {

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authenticated user. An user is required to get a record."})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	uc := usecase.NewEventTypeUseCase(h.EventTypeRepository, h.EventRepository)
	output, err := uc.GetEventType(c.Request.Context(), id, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (h *EventTypeHandler) GetAll(c *gin.Context) {

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authenticated user. An user is required to get records."})
		return
	}

	uc := usecase.NewEventTypeUseCase(h.EventTypeRepository, h.EventRepository)
	outputList, err := uc.GetAllEventTypes(c.Request.Context(), userId.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, outputList)
}
